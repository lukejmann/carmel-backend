package db

import (
	"fmt"
	"log"
	"testing"

	"github.com/gagliardetto/solana-go"
)

func TestGetMintCollections(t *testing.T) {

	testMint := solana.MustPublicKeyFromBase58("ys2ZiXC2tHtG6Eim67jEN169w5oqhNznsiJiYrdyoeY")

	fmt.Println("Testing mint:", testMint.String())

	collections, err := GetMintCollections(testMint)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("collections:", len(collections))
}
