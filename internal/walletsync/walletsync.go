package walletsync

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/getsentry/sentry-go"
	"github.com/lukejmann/carmel-backend/internal/constants"
	"github.com/lukejmann/carmel-backend/internal/db"
	"github.com/lukejmann/carmel-backend/internal/process"
	ah "github.com/lukejmann/carmel-backend/pkg/auctionhouse"
)

var rpcClient *rpc.Client

var MagicEdenProgram = solana.MustPublicKeyFromBase58("M2mx93ekt1fmXSVkTrUL9xVFHkmME8HTUi5Cyc5aF7K")
var HausProgram = solana.MustPublicKeyFromBase58("hausS13jsjafwWwGqZTUQRmWyvyxn9EQpqMwV1PBBmk")

func init() {
	endpoint := constants.RPCEndPoint
	rpcClient = rpc.New(endpoint)

	solana.RegisterInstructionDecoder(MagicEdenProgram, ah.RegistryDecodeInstruction)
	solana.RegisterInstructionDecoder(HausProgram, ah.RegistryDecodeInstruction)

	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://726714e0074349859e56c0895a01ea00@o985688.ingest.sentry.io/6417112",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
}

func SyncWalletToSlot(wallet solana.PublicKey, slot int) (err error) {
	defer sentry.Recover()

	fmt.Printf("Traversing wallet: %s\n", wallet)
	var BATCH_SIZE = 10
	var LOOP_LIMIT = 100

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
	walletObj, err := db.GetWallet(wallet)
	if err != nil {
		return err
	}
	var oldestProcessedSig *solana.Signature = nil
	if walletObj.OldestProcessedSig != "" {
		derivedSig, err := solana.SignatureFromBase58(walletObj.OldestProcessedSig)
		if err != nil {
			return err
		}
		fmt.Printf("Latest process sig: %s\n", walletObj.OldestProcessedSig)
		oldestProcessedSig = &derivedSig
	}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		var walletSigs []*rpc.TransactionSignature = nil
		var b = false
		for i := 0; i < LOOP_LIMIT; i++ {
			if b {
				wg.Done()
				return
			}
			if oldestProcessedSig == nil {
				walletSigs, err = rpcClient.GetSignaturesForAddressWithOpts(context.Background(), wallet, &rpc.GetSignaturesForAddressOpts{
					Limit: &BATCH_SIZE,
				})
				if err != nil {
					sentry.CaptureException(err)
				}
				if len(walletSigs) != 0 {
					err = db.SetWalletLatestProcessedSig(wallet, walletSigs[0].Signature.String())
					if err != nil {
						sentry.CaptureException(err)
					}
				}
			} else {
				walletSigs, err = rpcClient.GetSignaturesForAddressWithOpts(context.Background(), wallet, &rpc.GetSignaturesForAddressOpts{
					Limit:  &BATCH_SIZE,
					Before: *oldestProcessedSig,
				})
			}
			if err != nil {
				log.Fatal(err)
			}
			var wg2 sync.WaitGroup
			wg2.Add(len(walletSigs))
			for i, sig := range walletSigs {
				go func(i int, sig *rpc.TransactionSignature) {
					if b {
						wg2.Done()
						return
					}
					sigHasBeenProcessed, err := db.GetSigIsProcessedForWallet(wallet, sig.Signature)
					if err != nil {
						log.Fatal(err)
						sentry.CaptureException(err)
					}
					if !sigHasBeenProcessed {
						err = process.ProcessSignature(sig.Signature)
						if err != nil {
							log.Fatal(err)
							sentry.CaptureException(err)
						}
						err = db.SetSigProcessedForWallet(wallet, sig.Signature)
						if err != nil {
							log.Fatal(err)
							sentry.CaptureException(err)
						}
					} else {
						fmt.Printf("Sig already processed: %s\n", sig.Signature)
						b = true

					}
					fmt.Printf("Processed %d/%d\n", i+1, len(walletSigs))
					wg2.Done()
				}(i, sig)
			}
			wg2.Wait()
			if len(walletSigs) == 0 {
				break
			}
			err = db.SetWalletOldestProcessedSig(wallet, walletSigs[len(walletSigs)-1].Signature.String())
			if err != nil {
				sentry.CaptureException(err)
			}
			oldestProcessedSig = &walletSigs[len(walletSigs)-1].Signature
		}
		wg.Done()
	}()
	firebaseSyncedToSlot, err := db.GetWalletSyncedToSlot(wallet)
	if err != nil {
		return err
	}
	go func() {
		if firebaseSyncedToSlot+10 < slot {
			err = process.RefreshWalletData(wallet)
			if err != nil {
				sentry.CaptureException(err)
			}
		}
		wg.Done()
	}()
	err = db.SetWalletSyncedToSlot(wallet, slot)
	wg.Wait()

	return nil
}
