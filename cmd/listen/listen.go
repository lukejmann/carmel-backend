package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/lukejmann/carmel-backend/internal/constants"
	"google.golang.org/api/option"
)

var pubsubClient *pubsub.Client

func init() {
	ctx := context.Background()
	var authFile = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	pc, err := pubsub.NewClient(ctx, "projectc-59f56", option.WithCredentialsFile(authFile))
	if err != nil {
		log.Fatal(err)
	}
	pubsubClient = pc
}

func main() {
	err := Listen()
	if err != nil {
		log.Fatal(err)
	}
}

type ProcessSigMessage struct {
	Signature string `json:"signature"`
}

func Listen() (err error) {
	endpoint := rpc.MainNetBetaSerum_WS

	fmt.Println("In Listen")
	magicEdenClient, err := ws.Connect(context.Background(), endpoint)
	if err != nil {
		return err
	}
	defer magicEdenClient.Close()
	meSub, err := magicEdenClient.LogsSubscribeMentions(
		constants.MagicEdenProgram,
		rpc.CommitmentConfirmed,
	)
	if err != nil {
		return err
	}

	hausClient, err := ws.Connect(context.Background(), endpoint)
	if err != nil {
		return err
	}
	defer hausClient.Close()
	hausSub, err := hausClient.LogsSubscribeMentions(
		constants.HausProgram,
		rpc.CommitmentConfirmed,
	)
	if err != nil {
		return err
	}

	// metaplexClient, err := ws.Connect(context.Background(), endpoint)
	// if err != nil {
	// 	return err
	// }
	// defer metaplexClient.Close()
	// metaplexSub, err := metaplexClient.LogsSubscribeMentions(
	// 	constants.MetaplexProgram,
	// 	rpc.CommitmentConfirmed,
	// )
	// if err != nil {
	// 	return err
	// }

	topic := pubsubClient.Topic("processSig")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			meGot, err1 := meSub.Recv()
			if err1 != nil {
				fmt.Println("Error recieveing ME sig: ", err)
				continue
			}
			msg := ProcessSigMessage{
				Signature: meGot.Value.Signature.String(),
			}
			body, err := json.Marshal(msg)
			if err != nil {
				fmt.Println("Error marshalling ME sig: ", err)
			}
			spew.Dump(topic)
			r := topic.Publish(context.Background(), &pubsub.Message{
				Data: body,
			})
			select {
			case <-r.Ready():
				fmt.Println("Published ME message")
				sID, err := r.Get(context.Background())
				if err != nil {
					fmt.Println("Error publishing ME message: ", err)
				}
				fmt.Println("Published ME message with ID: ", sID)
			}
		}
	}()
	go func() {
		for {
			hausGot, err := hausSub.Recv()
			if err != nil {
				fmt.Println("Error recieveing haus sig: ", err)
				continue
			}
			msg := ProcessSigMessage{
				Signature: hausGot.Value.Signature.String(),
			}
			body, err := json.Marshal(msg)
			if err != nil {
				fmt.Println("Error marshalling haus sig: ", err)
			}
			r := topic.Publish(context.Background(), &pubsub.Message{
				Data: body,
			})
			select {
			case <-r.Ready():
				fmt.Println("Published haus message")
				sID, err := r.Get(context.Background())
				if err != nil {
					fmt.Println("Error publishing haus message: ", err)
				}
				fmt.Println("Published haus message with ID: ", sID)
			}
		}
	}()
	// go func() {
	// 	for {
	// 		metaplexGot, err := metaplexSub.Recv()
	// 		if err != nil {
	// 			fmt.Println("Error recieveing metaplex sig: ", err)
	// 			continue
	// 		}
	// 		msg := ProcessSigMessage{
	// 			Signature: metaplexGot.Value.Signature.String(),
	// 		}
	// 		body, err := json.Marshal(msg)
	// 		if err != nil {
	// 			fmt.Println("Error marshalling metaplex sig: ", err)
	// 		}
	// 		r := topic.Publish(context.Background(), &pubsub.Message{
	// 			Data: body,
	// 		})
	// 		select {
	// 		case <-r.Ready():
	// 			fmt.Println("Published metaplex message")
	// 			sID, err := r.Get(context.Background())
	// 			if err != nil {
	// 				fmt.Println("Error publishing metaplex message: ", err)
	// 			}
	// 			fmt.Println("Published metaplex message with ID: ", sID)
	// 		}
	// 	}
	// }()
	wg.Wait()

	return nil
}
