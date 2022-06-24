package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/lukejmann/carmel-backend/internal/process"
)

func main() {
	http.HandleFunc("/", ProcessSigListen)
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

type ProcessSigMessage struct {
	Signature string `json:"signature"`
}

func ProcessSigListen(w http.ResponseWriter, r *http.Request) {
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

	var processSigMessage ProcessSigMessage
	err = json.Unmarshal(m.Message.Data, &processSigMessage)
	if err != nil {
		msg := fmt.Errorf("Error decoding message: %s", err)
		log.Println(msg)
		return
	}
	log.Printf("ProcessSigListen signature: %s!\n", processSigMessage.Signature)
	pubkey, err := solana.SignatureFromBase58(processSigMessage.Signature)
	if err != nil {
		msg := fmt.Errorf("invalid signature: %s (%v)", processSigMessage.Signature, err)
		log.Println(msg)
		return
	}
	go process.ProcessSignature(pubkey)
}
