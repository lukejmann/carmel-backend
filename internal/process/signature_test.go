package process

import (
	"fmt"
	"log"
	"testing"

	"github.com/gagliardetto/solana-go"
)

func TestProcessSig(t *testing.T) {
	{
		testSig := solana.MustSignatureFromBase58("651biG67Dj3Uu2Qchauh8LRjWchqPBbFSiaNxYbQ7Yxio4jLYmm3yFJdhxzJdvFYL6aQAWUKLiJpkqT6BKmb8Gtc")
		fmt.Println("Processing sig:", testSig)

		err := ProcessSignature(testSig)
		if err != nil {
			log.Fatal(err)
		}
	}

}
