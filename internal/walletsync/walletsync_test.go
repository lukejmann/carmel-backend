package walletsync

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"

	"github.com/gagliardetto/solana-go"
)

func TestWalletTraverse(t *testing.T) {
	{
		var wg sync.WaitGroup
		wg.Add(1)
		shaq := solana.MustPublicKeyFromBase58("gacMrsrxNisAhCfgsUAVbwmTC3w9nJB6NychLAnTQFv")
		out, err := rpcClient.GetSlot(
			context.TODO(),
			"max",
		)
		if err != nil {
			panic(err)
		}
		go func() {
			err := SyncWalletToSlot(shaq, int(out))
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("done")
				wg.Done()
			}
		}()
		fmt.Println("End Listen?")
		wg.Wait()

	}
}
