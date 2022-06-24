package process

import (
	"context"
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/lukejmann/carmel-backend/internal/constants"
	"github.com/lukejmann/carmel-backend/internal/db"
	ah "github.com/lukejmann/carmel-backend/pkg/auctionhouse"
	"github.com/portto/solana-go-sdk/program/tokenprog"
)

func ProcessSignature(sig solana.Signature) error {
	alreadyProcessed, err := db.GetSigIsProcessed(sig)
	if alreadyProcessed {
		return nil
	}
	out, err := rpcClient.GetTransaction(
		context.TODO(),
		sig,
		&rpc.GetTransactionOpts{
			Encoding:   solana.EncodingBase64,
			Commitment: rpc.CommitmentConfirmed,
		},
	)
	if err != nil {
		return fmt.Errorf("error getting transaction: %v", err)
	}
	if out == nil {
		return fmt.Errorf("failed to get transaction2: %v", err)
	}
	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(out.Transaction.GetBinary()))
	if err != nil {
		return fmt.Errorf("failed to decode tx2: %v", err)
	}
	isAHProgram := false
	isTokenProgram := false
	// isCM1Program := false
	// isCM2Program := false
	// userIsSigner := false
	var signers = make([]string, 0)
	for _, signer := range tx.Message.Signers() {
		signers = append(signers, signer.String())
	}
	// if !userIsSigner {
	// 	fmt.Printf("Skipping tx: %s\n", txSig.Signature)
	// 	return nil
	// }
	for _, inst := range tx.Message.Instructions {
		progKey, err := tx.ResolveProgramIDIndex(inst.ProgramIDIndex)
		if err == nil {
			decodedInstruction, err := solana.DecodeInstruction(
				progKey,
				inst.ResolveInstructionAccounts(&tx.Message),
				inst.Data,
			)
			if err != nil {
				continue
			}
			_, ok := decodedInstruction.(*ah.Instruction)
			if ok {
				isAHProgram = true
			}

			_, ok = decodedInstruction.(*token.Instruction)
			if ok {
				isTokenProgram = true
			}
		} else {
			fmt.Printf("cannot ResolveProgramIDIndex: %s\n", err)
		}
	}
	progKey, err := tx.ResolveProgramIDIndex(tx.Message.Instructions[0].ProgramIDIndex)
	if err != nil {
		return fmt.Errorf("failed to resolve program id: %v", err)
	}
	if isAHProgram {
		isME := programIsME(progKey)
		err = ProcessMustAHSignature(tx, isME, out.Meta.LogMessages, int64(*out.BlockTime), signers, true)
		if err != nil {
			return err
		}
	}
	if isTokenProgram {
		// fmt.Printf("isTokenProgram\n")
		// spew.Dump(out)
		err = ProcessMustTokenSignature(tx, int64(*out.BlockTime), signers, true)
		if err != nil {
			return err
		}
	}
	// if isCM1Program {
	// 	fmt.Printf("isCM1Program\n")
	// 	panic("not implemented")
	// }
	// if isCM2Program {
	// 	fmt.Printf("isCM2Program\n")
	// 	panic("not implemented")
	// }
	err = db.SetSigProcessed(sig)
	if err != nil {
		return fmt.Errorf("failed to set sig processed: %v", err)
	}

	return nil
}

func ProcessMustAHSignature(tx *solana.Transaction, isME bool, logs []string, timestamp int64, signers []string, pushToDB bool) error {
	var buyInst *ah.Buy
	var sellInst *ah.Sell
	var executeInst *ah.ExecuteSale
	for _, inst := range tx.Message.Instructions {
		progKey, err := tx.ResolveProgramIDIndex(inst.ProgramIDIndex)
		if err != nil {
			panic(err)
		}
		decodedInstruction, err := solana.DecodeInstruction(
			progKey,
			inst.ResolveInstructionAccounts(&tx.Message),
			inst.Data,
		)
		if err != nil {
			continue
		}
		ahInst, ok := decodedInstruction.(*ah.Instruction)
		if ok {
			pExecuteInst, ok := ahInst.Impl.(*ah.ExecuteSale)
			if ok {
				executeInst = pExecuteInst
			}
			pBuyInst, ok := ahInst.Impl.(*ah.Buy)
			if ok {
				buyInst = pBuyInst
			}
			pSellInst, ok := ahInst.Impl.(*ah.Sell)
			if ok {
				sellInst = pSellInst
			}
		}
	}

	// MakeListing action
	if sellInst != nil && executeInst == nil {
		var user = sellInst.GetWalletAccount().PublicKey
		var price = sellInst.BuyerPrice
		var paymentMint string = SystemProgramID.String()
		var assetMint string = ""
		if isME {
			priceFromLogs, err := PriceFromLogs(logs)
			if err != nil {
				return err
			}
			price = priceFromLogs
			paymentMint = sellInst.AccountMetaSlice[1].PublicKey.String()
			assetMint = sellInst.AccountMetaSlice[4].PublicKey.String()
		} else {
			try, err := rpcClient.GetAccountInfoWithOpts(
				context.TODO(),
				sellInst.GetTokenAccountAccount().PublicKey,
				&rpc.GetAccountInfoOpts{
					Encoding: solana.EncodingBase64,
				},
			)
			if err != nil {
				return fmt.Errorf("failed to get transaction: %v", err)
			}

			data := try.Value.Data.GetBinary()
			got, err := tokenprog.TokenAccountFromData(data)
			if err != nil {
				return fmt.Errorf("failed to decode token account: %v", err)
			}
			assetMint = got.Mint.String()
		}
		makeListingAction := db.MakeListingAction{
			Type:        "makeListing",
			Wallet:      user.String(),
			AssetMint:   assetMint,
			AssetSize:   int64(1),
			PaymentMint: paymentMint,
			PaymentSize: int64(*price),
			Time:        timestamp,
		}
		if pushToDB {
			if err := db.AppendMakeListingAction(&makeListingAction); err != nil {
				return fmt.Errorf("failed to append make listing action: %v", err)
			}
			err := ProcessMint(solana.MustPublicKeyFromBase58(assetMint))
			if err != nil {
				return fmt.Errorf("failed to process mint for collections: %v", err)
			}
			err = processWallet(user)
			if err != nil {
				return fmt.Errorf("failed to shallow process wallet: %v", err)
			}
		} else {
			fmt.Printf("Make Listing Action: %+v\n", makeListingAction)
		}
	}

	// MakeOffer action
	if buyInst != nil && executeInst == nil {
		var price = buyInst.BuyerPrice
		var paymentMint string = SystemProgramID.String()
		var assetMint string = ""
		if isME {
			priceFromLogs, err := PriceFromLogs(logs)
			if err != nil {
				return err
			}
			price = priceFromLogs
			paymentMint = buyInst.AccountMetaSlice[1].PublicKey.String()
			assetMint = buyInst.AccountMetaSlice[2].PublicKey.String()
		} else {
			try, err := rpcClient.GetAccountInfoWithOpts(
				context.TODO(),
				buyInst.GetTokenAccountAccount().PublicKey,
				&rpc.GetAccountInfoOpts{
					Encoding: solana.EncodingBase64,
				},
			)
			fmt.Printf("processAHSell tokenAccountPubkey: %+v\n", buyInst.GetTokenAccountAccount().PublicKey.String())
			if err != nil {
				return fmt.Errorf("failed to get transaction: %v", err)
			}
			fmt.Printf("processAH Sell2\n")

			data := try.Value.Data.GetBinary()
			got, err := tokenprog.TokenAccountFromData(data)
			if err != nil {
				return fmt.Errorf("failed to decode token account: %v", err)
			}
			assetMint = got.Mint.String()
		}
		makeOfferAction := db.MakeOfferAction{
			Type:        "makeOffer",
			Wallet:      buyInst.GetWalletAccount().PublicKey.String(),
			AssetMint:   assetMint,
			AssetSize:   int64(1),
			PaymentMint: paymentMint,
			PaymentSize: int64(*price),
			Time:        timestamp,
		}
		if pushToDB {
			if err := db.AppendMakeOfferAction(&makeOfferAction); err != nil {
				return fmt.Errorf("failed to append make offer action: %v", err)
			}
			err := ProcessMint(solana.MustPublicKeyFromBase58(assetMint))
			if err != nil {
				return fmt.Errorf("failed to process mint for collections: %v", err)
			}
			err = processWallet(buyInst.GetWalletAccount().PublicKey)
			if err != nil {
				return fmt.Errorf("failed to shallow process wallet: %v", err)
			}
		} else {
			fmt.Printf("Make Offer Action: %+v\n", makeOfferAction)
		}
	}

	// Purchase action
	if executeInst != nil {
		var price = executeInst.BuyerPrice
		var buyer = executeInst.GetBuyerAccount().PublicKey
		var seller = executeInst.GetSellerAccount().PublicKey
		var paymentMint string = SystemProgramID.String()
		var assetMint string = ""
		if isME {
			priceFromLogs, err := PriceFromLogs(logs)
			if err != nil {
				return err
			}
			price = priceFromLogs
			paymentMint = executeInst.AccountMetaSlice[2].PublicKey.String()
			assetMint = executeInst.AccountMetaSlice[4].PublicKey.String()

		} else {
			tokenAccount := executeInst.GetTokenAccountAccount()
			if tokenAccount == nil {
				return fmt.Errorf("token account not found")
			}
			mint, err := MintFromTokenAccount(tokenAccount.PublicKey)
			if err != nil {
				return err
			}
			assetMint = mint.String()
		}
		var executor = buyer
		for _, signer := range signers {
			if signer == seller.String() {
				executor = seller
			}
		}
		purchaseAction := db.PurchaseAction{
			Type:        "purchase",
			Wallet:      executor.String(),
			Seller:      seller.String(),
			Buyer:       buyer.String(),
			AssetMint:   assetMint,
			AssetSize:   int64(1),
			PaymentMint: paymentMint,
			PaymentSize: int64(*price),
			Time:        timestamp,
		}

		if pushToDB {
			if err := db.AppendPurchaseAction(&purchaseAction); err != nil {
				return fmt.Errorf("failed to append purchase action: %v", err)
			}
			err := ProcessMint(solana.MustPublicKeyFromBase58(assetMint))
			if err != nil {
				return fmt.Errorf("failed to process asset mint for collections: %v", err)
			}
			err = processWallet(buyer)
			if err != nil {
				return fmt.Errorf("failed to shallow process wallet: %v", err)
			}
			err = processWallet(seller)
			if err != nil {
				return fmt.Errorf("failed to shallow process wallet: %v", err)
			}
		} else {
			fmt.Printf("Purchase Action: %+v\n", purchaseAction)
		}
		err := ProcessMint(solana.MustPublicKeyFromBase58(paymentMint))
		if err != nil {
			return fmt.Errorf("failed to process payment mint for collections: %v", err)
		}
	}
	return nil
}

func ProcessMustTokenSignature(tx *solana.Transaction, timestamp int64, signers []string, pushToDB bool) error {
	var transferInst *token.TransferChecked
	var mintInst *token.MintTo
	var mintCheckedInst *token.MintToChecked
	// var sellInst *ah.Sell
	// var executeInst *ah.ExecuteSale
	for _, inst := range tx.Message.Instructions {
		progKey, err := tx.ResolveProgramIDIndex(inst.ProgramIDIndex)
		if err != nil {
			panic(err)
		}
		decodedInstruction, err := solana.DecodeInstruction(
			progKey,
			inst.ResolveInstructionAccounts(&tx.Message),
			inst.Data,
		)
		if err != nil {
			continue
		}
		tokenInst, ok := decodedInstruction.(*token.Instruction)
		if ok {
			pTransferInst, ok := tokenInst.Impl.(*token.TransferChecked)
			if ok {
				transferInst = pTransferInst
			}

			pMintInst, ok := tokenInst.Impl.(*token.MintTo)
			if ok {
				mintInst = pMintInst
			}

			pMintCheckedInst, ok := tokenInst.Impl.(*token.MintToChecked)
			if ok {
				mintCheckedInst = pMintCheckedInst
			}
		}
	}
	if transferInst != nil {
		var source = transferInst.GetSourceAccount().PublicKey.String()
		var dest = transferInst.GetDestinationAccount().PublicKey.String()
		pMint, err := MintFromTokenAccount(transferInst.GetSourceAccount().PublicKey)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				fmt.Printf("failed to get mint from token account %v: %v\n", source, err)
				return nil
			} else {
				return fmt.Errorf("failed to get mint from token account %v: %v", source, err)
			}
		}
		mintIsNFT, err := MintIsNFT(*pMint)
		if err != nil {
			return fmt.Errorf("failed to get mint is nft: %v", err)
		}
		mint := pMint.String()
		var amount = int64(*transferInst.Amount)
		mintIsNFT = mintIsNFT && amount == 1
		if mintIsNFT {
			fmt.Printf("transfer of %d %s from %s to %s\n", amount, mint, source, dest)
			sender, err := OwnerFromTokenAccount(transferInst.GetSourceAccount().PublicKey)
			if err != nil {
				return fmt.Errorf("failed to get sender from token account %v: %v", source, err)
			}
			receiver, err := OwnerFromTokenAccount(transferInst.GetDestinationAccount().PublicKey)
			if err != nil {
				return fmt.Errorf("failed to get receiver from token account %v: %v", dest, err)
			}
			transferAction := db.TransferAction{
				Type:      "transfer",
				Wallet:    sender.String(),
				Sender:    sender.String(),
				Receiver:  receiver.String(),
				AssetMint: mint,
				AssetSize: amount,
				Time:      timestamp,
			}

			if pushToDB {
				if err := db.AppendTransferAction(&transferAction); err != nil {
					return fmt.Errorf("failed to append purchase action: %v", err)
				}
				err = ProcessMint(solana.MustPublicKeyFromBase58(mint))
				if err != nil {
					return fmt.Errorf("failed to process mint for collections: %v", err)
				}
				err = processWallet(*sender)
				if err != nil {
					return fmt.Errorf("failed to shallow process wallet: %v", err)
				}
				err = processWallet(*receiver)
				if err != nil {
					return fmt.Errorf("failed to shallow process wallet: %v", err)
				}
			} else {
				fmt.Printf("Transfer Action: \n")
				spew.Dump(transferAction)
			}
		} else {
			fmt.Printf("mint is not NFT\n")
		}
	}
	if mintInst != nil {
		mint := mintInst.GetMintAccount().PublicKey
		dest := mintInst.GetDestinationAccount().PublicKey
		amount := int64(*mintInst.Amount)
		fmt.Printf("mint of %d %s to %s\n", amount, mint, dest)

		mintAction := db.MintAction{
			Type:      "mint",
			Wallet:    dest.String(),
			AssetMint: mint.String(),
			AssetSize: amount,
			Time:      timestamp,
		}
		if err := db.AppendMintAction(&mintAction); err != nil {
			return fmt.Errorf("failed to append mint action: %v", err)
		}

		err := ProcessMint(solana.MustPublicKeyFromBase58(mint.String()))
		if err != nil {
			return fmt.Errorf("failed to process mint for collections: %v", err)
		}

		err = processWallet(dest)
		if err != nil {
			return fmt.Errorf("failed to shallow process wallet: %v", err)
		}
	}
	if mintCheckedInst != nil {
		mint := mintCheckedInst.GetMintAccount().PublicKey
		dest := mintCheckedInst.GetDestinationAccount().PublicKey
		amount := int64(*mintCheckedInst.Amount)
		fmt.Printf("mint checked of %d %s to %s\n", amount, mint, dest)
		mintAction := db.MintAction{
			Type:      "mint",
			Wallet:    dest.String(),
			AssetMint: mint.String(),
			AssetSize: amount,
			Time:      timestamp,
		}
		if pushToDB {
			if err := db.AppendMintAction(&mintAction); err != nil {
				return fmt.Errorf("failed to append mint action: %v", err)
			}

			err := ProcessMint(solana.MustPublicKeyFromBase58(mint.String()))
			if err != nil {
				return fmt.Errorf("failed to process mint for collections: %v", err)
			}
			err = processWallet(dest)
			if err != nil {
				return fmt.Errorf("failed to shallow process wallet: %v", err)
			}
		} else {
			fmt.Printf("Mint Action: \n")
			spew.Dump(mintAction)
		}
	}

	return nil
}

func programIsME(prog solana.PublicKey) bool {
	// fmt.Printf("prog: %s\n", prog)
	return prog == constants.MagicEdenProgram
}

func processWallet(wallet solana.PublicKey) error {
	exists, err := db.WalletExists(wallet)
	if err != nil {
		return err
	}
	if !exists {
		err := db.CreateEmptyWalletDoc(wallet)
		if err != nil {
			return err
		}
	}
	return nil
}
