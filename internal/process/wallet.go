package process

import (
	"context"
	"crypto/sha256"
	"fmt"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/lukejmann/carmel-backend/internal/constants"
	"github.com/lukejmann/carmel-backend/internal/db"
)

func RefreshWalletData(wallet solana.PublicKey) (err error) {
	out, err := rpcClient.GetTokenAccountsByOwner(
		context.TODO(),
		wallet,
		&rpc.GetTokenAccountsConfig{
			ProgramId: &solana.TokenProgramID,
		},
		&rpc.GetTokenAccountsOpts{
			Encoding: solana.EncodingBase64Zstd,
		},
	)
	if err != nil {
		panic(err)
	}
	ownedTokens := make([]db.OwnedTokens, 0)
	{
		tokenAccounts := make([]token.Account, 0)
		for _, rawAccount := range out.Value {
			var tokAcc token.Account

			data := rawAccount.Account.Data.GetBinary()
			dec := bin.NewBinDecoder(data)
			err := dec.Decode(&tokAcc)
			if err != nil {
				panic(err)
			}
			tokenAccounts = append(tokenAccounts, tokAcc)
			if tokAcc.Amount > 0 {
				ownedToken := db.OwnedTokens{
					Mint: tokAcc.Mint.String(),
					Size: int(tokAcc.Amount),
				}
				err := ProcessMint(tokAcc.Mint)
				if err != nil {
					return fmt.Errorf("failed to process mint: %v", err)
				}
				ownedTokens = append(ownedTokens, ownedToken)
			}
		}
	}
	if len(ownedTokens) > 0 {
		err = db.SetWalletOwnedTokens(wallet, ownedTokens)
		if err != nil {
			return fmt.Errorf("failed to set wallet owned tokens: %v", err)
		}
	}
	return nil
}

func GetOwnedDomainsForWallet(wallet solana.PublicKey) (ownedDomains []string, err error) {
	out, err := rpcClient.GetProgramAccountsWithOpts(
		context.TODO(),
		constants.NameServicePubkey,
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{
				{
					Memcmp: &rpc.RPCFilterMemcmp{
						Offset: 32,
						Bytes:  wallet.Bytes(),
					},
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}

	keys := make([]solana.PublicKey, 0)
	for _, rawAccount := range out {
		keys = append(keys, rawAccount.Pubkey)
	}

	centralState, bump, err := solana.FindProgramAddress([][]byte{constants.NameServicePubkey.Bytes()}, constants.NameServicePubkey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("central state, bump: %v, %v\n", centralState, bump)
	hashPrefix := "SPL Name Service"
	hashes := make([]byte, 0)
	for _, key := range keys {
		input := hashPrefix + key.String()
		h := sha256.New()
		h.Write([]byte(input))
		hashes = append(hashes, h.Sum(nil)...)
	}

	// TODO: finish implementation
	return nil, nil
}
