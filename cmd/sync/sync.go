package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/lukejmann/carmel-backend/internal/walletsync"
)

func main() {
	http.HandleFunc("/", SyncSub)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

type SyncToSlotMessage struct {
	Wallet     string `json:"wallet"`
	SyncToSlot int    `json:"syncToSlot"`
}

func SyncSub(w http.ResponseWriter, r *http.Request) {
	var m PubSubMessage
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		return
	}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		return
	}

	log.Printf("SigSub subscription: %s!\n", m.Subscription)
	var syncToSlotMsg SyncToSlotMessage
	err = json.Unmarshal(m.Message.Data, &syncToSlotMsg)
	if err != nil {
		msg := fmt.Errorf("Error decoding message: %s", err)
		log.Println(msg)
		return
	}
	log.Printf("SigSub pubkey: %s!\n", syncToSlotMsg.Wallet)
	log.Printf("SigSub slot: %v!\n", syncToSlotMsg.SyncToSlot)
	pubkey, err := solana.PublicKeyFromBase58(syncToSlotMsg.Wallet)
	if err != nil {
		msg := fmt.Errorf("invalid pubkey: %s (%v)", syncToSlotMsg.Wallet, err)
		log.Println(msg)
		return
	}
	go walletsync.SyncWalletToSlot(pubkey, int(syncToSlotMsg.SyncToSlot))
}
