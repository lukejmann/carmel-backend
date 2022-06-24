package process

import (
	"context"
	"fmt"
	"log"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/lukejmann/carmel-backend/internal/db"
	cm1 "github.com/lukejmann/carmel-backend/pkg/cm1"
)

func UpdateOnChainCollectionMints(collectionKey solana.PublicKey) error {
	var offset uint64 = 0
	var length uint64 = 412
	accounts, err := rpcClient.GetProgramAccountsWithOpts(context.Background(), solana.TokenMetadataProgramID, &rpc.GetProgramAccountsOpts{
		Commitment: rpc.CommitmentConfirmed,
		Encoding:   solana.EncodingBase58,
		DataSlice: &rpc.DataSlice{
			Offset: &offset,
			Length: &length,
		},
		Filters: []rpc.RPCFilter{
			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: uint64(offset),
					Bytes:  collectionKey.Bytes(),
				},
			},
		},
	})
	fmt.Printf("accounts: %+v\n", accounts)
	if err != nil {
		return err
	}

	fmt.Printf("accounts: %+v\n", accounts)

	var collectionMints []solana.PublicKey = []solana.PublicKey{}

	for _, account := range accounts {
		metadataPubkey := account.Pubkey
		try, err := rpcClient.GetAccountInfoWithOpts(
			context.TODO(),
			metadataPubkey,
			&rpc.GetAccountInfoOpts{
				Encoding: solana.EncodingBase64,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to get transaction: %v", err)
		}

		data := try.Value.Data.GetBinary()
		metadata, err := MetadataDeserialize(data)
		mintPubkey := metadata.Mint
		collectionMints = append(collectionMints, mintPubkey)
	}
	err = db.ReplaceCollectionMints(collectionKey, collectionMints)
	if err != nil {
		return err
	}
	fmt.Printf("collectionMints: %v\n", collectionMints)
	return nil
}

func AddCollectionFromOnChainKey(collectionMint solana.PublicKey) error {
	metadataAccount, err := GetTokenMetaPubkey(collectionMint)
	if err != nil {
		log.Fatalf("faield to get metadata account, err: %v", err)
	}
	fmt.Println("collectionMint:", collectionMint)
	fmt.Println("metadataAccount:", metadataAccount)
	try, err := rpcClient.GetAccountInfoWithOpts(
		context.Background(),
		metadataAccount,
		&rpc.GetAccountInfoOpts{
			Encoding: solana.EncodingBase64,
		},
	)
	if err != nil {
		if err.Error() == "not found" {
			err = db.CreateEmptyOnChainCollection(collectionMint, "Untitled Collection", "Untitled Description", "")
			if err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("failed to get transaction: %v", err)
	}

	data := try.Value.Data.GetBinary()
	metadata, err := MetadataDeserialize(data)
	if err != nil {
		return fmt.Errorf("failed to deserialize mint metadata: %v", err)
	}

	exists, err := db.CollectionExists(collectionMint)
	if err != nil {
		return fmt.Errorf("failed to check if collection exists: %v", err)
	}
	fmt.Printf("exists: %v\n", exists)
	if !exists {
		extended, err := FetchExtendedMetadata(metadata.Data.Uri)
		fmt.Printf("collection extended: %v\n", extended)
		// spew.Dump(extended)
		if err != nil {
			fmt.Printf("failed to fetch extended metadata: %v\n", err)
			err = db.CreateEmptyOnChainCollection(collectionMint, metadata.Data.Name, "No Extended Metadata", "No Extended Metadata")
			if err != nil {
				return err
			}
			return nil
		}
		extendedDeref := *extended
		err = db.CreateEmptyOnChainCollection(collectionMint, extendedDeref.Name, extendedDeref.Description, extendedDeref.Image)
		if err != nil {
			return err
		}
	} else {

	}
	return nil
}

func UpdateCandyMachineCollectionMints(candyMachine solana.PublicKey) (collectionMints []solana.PublicKey, err error) {
	var limit int = 1000
	var before solana.Signature = solana.Signature{}
	var allSignatures []rpc.TransactionSignature = []rpc.TransactionSignature{}
	for i := 0; i < 100; i++ {
		accounts, err := rpcClient.GetSignaturesForAddressWithOpts(context.Background(), candyMachine, &rpc.GetSignaturesForAddressOpts{
			Limit:      &limit,
			Before:     before,
			Commitment: rpc.CommitmentConfirmed,
		})
		if err != nil {
			return nil, err
		}
		if len(accounts) == 0 {
			break
		}
		for _, account := range accounts {
			allSignatures = append(allSignatures, *account)
		}
		before = accounts[len(accounts)-1].Signature
	}

	for _, sig := range allSignatures {
		sigDetails, err := rpcClient.GetTransaction(context.Background(), sig.Signature, &rpc.GetTransactionOpts{
			Encoding: solana.EncodingBase64,
		})
		if err != nil {
			return nil, err
		}
		if sigDetails.Meta.Err != nil {
			continue
		}
		tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(sigDetails.Transaction.GetBinary()))
		if err != nil {
			return nil, fmt.Errorf("failed to decode tx2: %v", err)
		}
		insts := tx.Message.Instructions
		for _, rawInst := range insts {
			inst, err := cm1.DecodeInstruction(rawInst.ResolveInstructionAccounts(&tx.Message), rawInst.Data)
			if err != nil {
				// fmt.Printf("err at %v: %+v\n", i, err)
				continue
			}
			if inst.BaseVariant.TypeID.Equal(cm1.Instruction_MintNft.Bytes()) {
				mintInst := inst.Impl.(*cm1.MintNft)
				collectionMints = append(collectionMints, mintInst.GetMintAccount().PublicKey)
			}
		}
	}
	return collectionMints, err
}
