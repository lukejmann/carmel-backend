package db

import (
	"context"
	"fmt"
)

type MakeListingAction struct {
	ID          string `json:"id" firestore:"id"`
	Type        string `json:"type" firestore:"type"`
	Wallet      string `json:"wallet" firestore:"wallet"`
	AssetMint   string `json:"assetMint" firestore:"assetMint"`
	AssetSize   int64  `json:"assetSize" firestore:"assetSize"`
	PaymentMint string `json:"paymentMint" firestore:"paymentMint"`
	PaymentSize int64  `json:"paymentSize" firestore:"paymentSize"`
	Time        int64  `json:"time" firestore:"time"`
}

func AppendMakeListingAction(makeListingAction *MakeListingAction) error {
	if makeListingAction.Type != "makeListing" {
		return fmt.Errorf("makeListingAction.Type must be 'makeListing'")
	}
	collectionRef := firestoreClient.Collection(ActionsCollection)
	newMintRef := collectionRef.NewDoc()
	makeListingAction.ID = newMintRef.ID
	_, err := newMintRef.Set(context.Background(), makeListingAction)
	if err != nil {
		return err
	}
	return nil
}

type MakeOfferAction struct {
	ID          string `json:"id" firestore:"id"`
	Type        string `json:"type" firestore:"type"`
	Wallet      string `json:"wallet" firestore:"wallet"`
	AssetMint   string `json:"assetMint" firestore:"assetMint"`
	AssetSize   int64  `json:"assetSize" firestore:"assetSize"`
	PaymentMint string `json:"paymentMint" firestore:"paymentMint"`
	PaymentSize int64  `json:"paymentSize" firestore:"paymentSize"`
	Time        int64  `json:"time" firestore:"time"`
}

func AppendMakeOfferAction(makeOfferAction *MakeOfferAction) error {
	if makeOfferAction.Type != "makeOffer" {
		return fmt.Errorf("makeOfferAction.Type must be 'makeOffer'")
	}
	collectionRef := firestoreClient.Collection(ActionsCollection)
	newMintRef := collectionRef.NewDoc()
	makeOfferAction.ID = newMintRef.ID
	_, err := newMintRef.Set(context.Background(), makeOfferAction)
	if err != nil {
		return err
	}
	return nil
}

type PurchaseAction struct {
	ID          string `json:"id" firestore:"id"`
	Type        string `json:"type" firestore:"type"`
	Wallet      string `json:"wallet" firestore:"wallet"`
	Seller      string `json:"seller" firestore:"seller"`
	Buyer       string `json:"buyer" firestore:"buyer"`
	AssetMint   string `json:"assetMint" firestore:"assetMint"`
	AssetSize   int64  `json:"assetSize" firestore:"assetSize"`
	PaymentMint string `json:"paymentMint" firestore:"paymentMint"`
	PaymentSize int64  `json:"paymentSize" firestore:"paymentSize"`
	Time        int64  `json:"time" firestore:"time"`
}

func AppendPurchaseAction(purchaseAction *PurchaseAction) error {
	if purchaseAction.Type != "purchase" {
		return fmt.Errorf("purchaseAction.Type must be 'purchase'")
	}
	collectionRef := firestoreClient.Collection(ActionsCollection)
	newMintRef := collectionRef.NewDoc()
	purchaseAction.ID = newMintRef.ID
	_, err := newMintRef.Set(context.Background(), purchaseAction)
	if err != nil {
		return err
	}
	return nil
}

type CancelListingAction struct {
	ID          string `json:"id" firestore:"id"`
	Type        string `json:"type" firestore:"type"`
	Wallet      string `json:"wallet" firestore:"wallet"`
	AssetMint   string `json:"assetMint" firestore:"assetMint"`
	AssetSize   int64  `json:"assetSize" firestore:"assetSize"`
	PaymentMint string `json:"paymentMint" firestore:"paymentMint"`
	PaymentSize int64  `json:"paymentSize" firestore:"paymentSize"`
	Time        int64  `json:"time" firestore:"time"`
}

func AppendCancelListingAction(cancelListingAction *CancelListingAction) error {
	if cancelListingAction.Type != "cancelListing" {
		return fmt.Errorf("cancelListingAction.Type must be 'cancelListing'")
	}
	collectionRef := firestoreClient.Collection(ActionsCollection)
	newMintRef := collectionRef.NewDoc()
	cancelListingAction.ID = newMintRef.ID
	_, err := newMintRef.Set(context.Background(), cancelListingAction)
	if err != nil {
		return err
	}
	return nil
}

type TransferAction struct {
	ID        string `json:"id" firestore:"id"`
	Type      string `json:"type" firestore:"type"`
	Wallet    string `json:"wallet" firestore:"wallet"`
	Sender    string `json:"sender" firestore:"sender"`
	Receiver  string `json:"receiver" firestore:"receiver"`
	AssetMint string `json:"assetMint" firestore:"assetMint"`
	AssetSize int64  `json:"assetSize" firestore:"assetSize"`
	Time      int64  `json:"time" firestore:"time"`
}

func AppendTransferAction(transferAction *TransferAction) error {
	if transferAction.Type != "transfer" {
		return fmt.Errorf("transferAction.Type must be 'transfer'")
	}
	collectionRef := firestoreClient.Collection(ActionsCollection)
	newMintRef := collectionRef.NewDoc()
	transferAction.ID = newMintRef.ID
	_, err := newMintRef.Set(context.Background(), transferAction)
	if err != nil {
		return err
	}
	return nil
}

type MintAction struct {
	ID        string `json:"id" firestore:"id"`
	Type      string `json:"type" firestore:"type"`
	Wallet    string `json:"wallet" firestore:"wallet"`
	AssetMint string `json:"assetMint" firestore:"assetMint"`
	AssetSize int64  `json:"assetSize" firestore:"assetSize"`
	Time      int64  `json:"time" firestore:"time"`
}

func AppendMintAction(mintAction *MintAction) error {
	if mintAction.Type != "mint" {
		return fmt.Errorf("mintAction.Type must be 'mint'")
	}
	collectionRef := firestoreClient.Collection(ActionsCollection)
	newMintRef := collectionRef.NewDoc()
	mintAction.ID = newMintRef.ID
	_, err := newMintRef.Set(context.Background(), mintAction)
	if err != nil {
		return err
	}
	return nil
}
