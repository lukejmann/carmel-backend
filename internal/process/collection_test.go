package process

import (
	"fmt"
	"log"
	"testing"

	"github.com/gagliardetto/solana-go"
)

func TestUpdateOnChainCollectionMints(t *testing.T) {
	testCollection := solana.MustPublicKeyFromBase58("6XxjKYFbcndh2gDcsUrmZgVEsoDxXMnfsaGY6fpTJzNr")

	fmt.Println("Testing mint:", testCollection.String())

	err := AddCollectionFromOnChainKey(testCollection)
	if err != nil {
		log.Fatal(err)
	}
}
