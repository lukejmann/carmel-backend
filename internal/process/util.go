package process

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/lukejmann/carmel-backend/internal/db"
	cm1 "github.com/lukejmann/carmel-backend/pkg/cm1"
	cm2 "github.com/lukejmann/carmel-backend/pkg/cm2"
	"github.com/near/borsh-go"
	"github.com/portto/solana-go-sdk/program/tokenprog"
)

var (
	SystemProgramID                    = solana.MustPublicKeyFromBase58("11111111111111111111111111111111")
	ConfigProgramID                    = solana.MustPublicKeyFromBase58("Config1111111111111111111111111111111111111")
	StakeProgramID                     = solana.MustPublicKeyFromBase58("Stake11111111111111111111111111111111111111")
	VoteProgramID                      = solana.MustPublicKeyFromBase58("Vote111111111111111111111111111111111111111")
	BPFLoaderProgramID                 = solana.MustPublicKeyFromBase58("BPFLoader1111111111111111111111111111111111")
	Secp256k1ProgramID                 = solana.MustPublicKeyFromBase58("KeccakSecp256k11111111111111111111111111111")
	TokenProgramID                     = solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	MemoProgramID                      = solana.MustPublicKeyFromBase58("MemoSq4gqABAXKb96qnH8TysNcWxMyWCqXgDLGmfcHr")
	SPLAssociatedTokenAccountProgramID = solana.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL")
	SPLNameServiceProgramID            = solana.MustPublicKeyFromBase58("namesLPneVptA9Z5rqUDD9tMTWEJwofgaYwp8cawRkX")
	MetaplexTokenMetaProgramID         = solana.MustPublicKeyFromBase58("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")
	ComputeBudgetProgramID             = solana.MustPublicKeyFromBase58("ComputeBudget111111111111111111111111111111")
)

func GetTokenMetaPubkey(mint solana.PublicKey) (solana.PublicKey, error) {
	metadataAccount, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
		},
		MetaplexTokenMetaProgramID,
	)
	if err != nil {
		return solana.PublicKey{}, err
	}
	return metadataAccount, nil
}

func GetMasterEdition(mint solana.PublicKey) (solana.PublicKey, error) {
	msaterEdtion, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			[]byte("edition"),
		},
		MetaplexTokenMetaProgramID,
	)
	if err != nil {
		return solana.PublicKey{}, err
	}
	return msaterEdtion, nil
}

func GetEditionMark(mint solana.PublicKey, edition uint64) (solana.PublicKey, error) {
	editionNumber := edition / db.EDITION_MARKER_BIT_SIZE
	pubkey, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
			[]byte("edition"),
			[]byte(strconv.FormatUint(editionNumber, 10)),
		},
		MetaplexTokenMetaProgramID,
	)
	return pubkey, err
}

func FindCandyMachineFromMint(mint solana.PublicKey) (*solana.PublicKey, error) {
	var limit int = 1000
	var before solana.Signature = solana.Signature{}
	var allSignatures []rpc.TransactionSignature = []rpc.TransactionSignature{}
	for i := 0; i < 100; i++ {
		fmt.Printf("i: %+v\n", i)
		accounts, err := rpcClient.GetSignaturesForAddressWithOpts(context.Background(), mint, &rpc.GetSignaturesForAddressOpts{
			Limit:      &limit,
			Before:     before,
			Commitment: rpc.CommitmentConfirmed,
		})
		if err != nil {
			return nil, fmt.Errorf("error getting signatures for mint: %v", err)
		}
		if len(accounts) == 0 {
			break
		}
		for _, account := range accounts {
			allSignatures = append(allSignatures, *account)
		}
		before = accounts[len(accounts)-1].Signature
	}
	oldestSig := allSignatures[len(allSignatures)-1]
	sigDetails, err := rpcClient.GetTransaction(context.Background(), oldestSig.Signature, &rpc.GetTransactionOpts{
		Encoding: solana.EncodingBase64,
	})
	// spew.Dump(sigDetails)
	if err != nil {
		return nil, fmt.Errorf("error getting transaction details for mint %v: %v", mint, err)
	}
	// spew.Dump(sigDetails)
	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(sigDetails.Transaction.GetBinary()))
	if err != nil {
		return nil, fmt.Errorf("failed to decode tx2: %v", err)
	}
	insts := tx.Message.Instructions
	for _, rawInst := range insts {
		inst, err := cm1.DecodeInstruction(rawInst.ResolveInstructionAccounts(&tx.Message), rawInst.Data)
		if err != nil {
			continue
		}
		if inst.BaseVariant.TypeID.Equal(cm1.Instruction_MintNft.Bytes()) {
			mintInst := inst.Impl.(*cm1.MintNft)
			if mint.Equals(mintInst.GetMintAccount().PublicKey) {
				return &mintInst.GetCandyMachineAccount().PublicKey, nil
			}
		}

		inst2, err := cm2.DecodeInstruction(rawInst.ResolveInstructionAccounts(&tx.Message), rawInst.Data)
		if err != nil {
			continue
		}
		if inst2.BaseVariant.TypeID.Equal(cm2.Instruction_MintNft.Bytes()) {
			mintInst := inst2.Impl.(*cm2.MintNft)
			if mint.Equals(mintInst.GetMintAccount().PublicKey) {
				return &mintInst.GetCandyMachineAccount().PublicKey, nil
			}
		}
	}
	return nil, nil
}

func MetadataDeserialize(data []byte) (db.Metadata, error) {
	var metadata db.Metadata
	err := borsh.Deserialize(&metadata, data)
	if err != nil {
		return db.Metadata{}, fmt.Errorf("failed to deserialize data, err: %v", err)
	}
	// trim null byte
	metadata.Data.Name = strings.TrimRight(metadata.Data.Name, "\x00")
	metadata.Data.Symbol = strings.TrimRight(metadata.Data.Symbol, "\x00")
	metadata.Data.Uri = strings.TrimRight(metadata.Data.Uri, "\x00")
	return metadata, nil
}

func FetchExtendedMetadata(extendedURI string) (*db.MetadataExtended, error) {
	fmt.Printf("fetching extended metadata from %v\n", extendedURI)
	req, err := http.NewRequest(http.MethodGet, extendedURI, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("User-Agent", "space-cadet")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get extended metadata: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read extended metadata: %v", err)
	}

	extended := db.MetadataExtended{}

	jsonErr := json.Unmarshal(body, &extended)
	if jsonErr != nil {
		return nil, fmt.Errorf("failed to unmarshal extended metadata: %v", jsonErr)
	}
	// fmt.Printf("extended metadata: %v\n", extended)
	// spew.Dump(extended)

	return &extended, nil
}

func MintIsNFT(mintPubkey solana.PublicKey) (bool, error) {
	metadataAccount, err := GetTokenMetaPubkey(mintPubkey)
	if err != nil {
		log.Fatalf("faield to get metadata account, err: %v", err)
	}
	_, err = rpcClient.GetAccountInfoWithOpts(
		context.TODO(),
		metadataAccount,
		&rpc.GetAccountInfoOpts{
			Encoding: solana.EncodingBase64,
		},
	)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func MintFromTokenAccount(tokenAccount solana.PublicKey) (*solana.PublicKey, error) {
	try, err := rpcClient.GetAccountInfoWithOpts(
		context.TODO(),
		tokenAccount,
		&rpc.GetAccountInfoOpts{
			Encoding: solana.EncodingBase64,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get account %v: %v", tokenAccount, err)
	}
	data := try.Value.Data.GetBinary()
	got, err := tokenprog.TokenAccountFromData(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode token account: %v", err)
	}
	mintPubkey := solana.MustPublicKeyFromBase58(got.Mint.String())
	return &mintPubkey, nil
}

func PriceFromLogs(messages []string) (*uint64, error) {
	var sellLog MESellLog
	for _, log := range messages {
		if strings.Contains(log, "price") && len(log) > 12 {
			cut := log[12:]
			err := json.Unmarshal([]byte(cut), &sellLog)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal log: %v", err)
			}
		}
	}
	return &sellLog.Price, nil
}

func OwnerFromTokenAccount(tokenAccount solana.PublicKey) (*solana.PublicKey, error) {
	try, err := rpcClient.GetAccountInfoWithOpts(
		context.TODO(),
		tokenAccount,
		&rpc.GetAccountInfoOpts{
			Encoding: solana.EncodingBase64,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get account %v: %v", tokenAccount, err)
	}
	var account token.Account
	err = bin.NewBinDecoder(try.Value.Data.GetBinary()).Decode(&account)
	if err != nil {
		panic(err)
	}
	owner := account.Owner
	return &owner, nil
}

func GetTransactionFromSigStr(sigStr string) (*solana.Transaction, error) {
	sig, err := solana.SignatureFromBase58(sigStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode signature: %v", err)
	}
	txRes, err := rpcClient.GetTransaction(context.TODO(), sig, &rpc.GetTransactionOpts{
		Encoding: solana.EncodingBase64,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %v", err)
	}
	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(txRes.Transaction.GetBinary()))
	if err != nil {
		return nil, fmt.Errorf("failed to decode transaction: %v", err)
	}
	return tx, nil
}
