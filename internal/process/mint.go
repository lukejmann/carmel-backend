package process

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/lukejmann/carmel-backend/internal/constants"
	"github.com/lukejmann/carmel-backend/internal/db"
)

func ProcessMint(mintPubkey solana.PublicKey) error {
	fmt.Println("Adding mint:", mintPubkey)
	exists, err := db.MintExists(mintPubkey)
	if err != nil {
		return fmt.Errorf("failed to check if mint exists: %v", err)
	}
	if mintPubkey.String() == constants.SolBase58 {
		return nil
	}
	if exists {
		// fmt.Println("Mint already exists, skipping")
		return nil
	}
	metadataAccount, err := GetTokenMetaPubkey(mintPubkey)
	if err != nil {
		return fmt.Errorf("failed to get metadata account: %v", err)
	}
	try, err := rpcClient.GetAccountInfoWithOpts(
		context.TODO(),
		metadataAccount,
		&rpc.GetAccountInfoOpts{
			Encoding: solana.EncodingBase64,
		},
	)
	if err != nil {
		if err.Error() == "not found" {
			fmt.Println("Mint does not have metadata, skipping")
			return nil
		}
		return fmt.Errorf("failed to get transaction: %v", err)
	}

	data := try.Value.Data.GetBinary()
	// spew.Dump(data)
	metadata, err := MetadataDeserialize(data)
	if err != nil {
		return fmt.Errorf("failed to deserialize mint metadata: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, metadata.Data.Uri, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "space-cadet")

	if metadata.Collection.Verified != 0 {
		fmt.Printf("metadata.Collection.Key: %v\n", metadata.Collection.Key)
		exists, err := db.CollectionExists(metadata.Collection.Key)
		fmt.Printf("exists: %v\n", exists)
		if err != nil {
			return fmt.Errorf("failed to check if collection exists: %v", err)
		}
		if !exists {
			err = AddCollectionFromOnChainKey(metadata.Collection.Key)
			if err != nil {
				return fmt.Errorf("failed to add collection: %v", err)
			}
		} else {

		}
		extended, err := FetchExtendedMetadata(metadata.Data.Uri)
		if err == nil && extended != nil {
			err = db.AddMintWithMetadataExtended(mintPubkey, metadata, extended)
			if err != nil {
				return err
			}
		} else {
			err = db.AddMintWithMetadata(mintPubkey, metadata)
			if err != nil {
				return err
			}
		}
		err = db.AddMintToCollection(metadata.Collection.Key, mintPubkey)
		if err != nil {
			return fmt.Errorf("failed to add mint to on-chain collection: %v", err)
		}
		return nil
	}

	candyMachine, err := FindCandyMachineFromMint(mintPubkey)
	if err != nil {
		return fmt.Errorf("err finding candy machine from mint: %v", err)
	}
	if candyMachine != nil {
		fmt.Printf("candy machine: %v\n", candyMachine)
		exists, err := db.CollectionExists(*candyMachine)
		if err != nil {
			return fmt.Errorf("failed to check if collection exists: %v", err)
		}
		fmt.Printf("candy machine exists: %v\n", exists)
		if !exists {
			extended, err := FetchExtendedMetadata(metadata.Data.Uri)
			if err == nil && extended != nil && extended.Collection != nil {
				err = db.CreateEmptyCandyMachineCollection(*candyMachine, extended.Collection.Name)
				if err != nil {
					return err
				}
				// return nil
			}
			if err != nil {
				fmt.Printf("failed to fetch extended metadata: %v", err)
			}
			fmt.Printf("metadata collection empty. uri: %s\n", metadata.Data.Uri)
			err = db.CreateEmptyCandyMachineCollection(*candyMachine, fmt.Sprintf("Candy Machine %v", candyMachine.String()))
			if err != nil {
				return err
			}
		}
		extended, err := FetchExtendedMetadata(metadata.Data.Uri)
		if err == nil && extended != nil {
			err = db.AddMintWithMetadataExtended(mintPubkey, metadata, extended)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("failed to fetch extended metadata: %v", err)
			err = db.AddMintWithMetadata(mintPubkey, metadata)
			if err != nil {
				return err
			}
		}
		err = db.AddMintToCollection(*candyMachine, mintPubkey)
		if err != nil {
			return fmt.Errorf("failed to add mint to candy machine collection: %v", err)
		}
	} else {
		fmt.Printf("failed to find candy machine from mint: %v (attempting add without collection\n", err)
		extended, err := FetchExtendedMetadata(metadata.Data.Uri)
		if err == nil && extended != nil {
			err = db.AddMintWithMetadataExtended(mintPubkey, metadata, extended)
			if err != nil {
				return err
			}
		} else {
			err = db.AddMintWithMetadata(mintPubkey, metadata)
			if err != nil {
				return err
			}

		}
		return nil
	}

	return nil
}
