package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var rpcClient *rpc.Client
var bqClient *bigquery.Client
var firestoreClient *firestore.Client

var projectID = "projectc-59f56"

var CollectionCollection = "collections4"
var MintsCollection = "mints4"
var ActionsCollection = "actions4"
var WalletsCollection = "wallets4"
var WalletsSigsCollection = "walletsSigs4"
var ProcessedSigsCollection = "processedSigsCollection4"

var OrdersCollection = "ordersB6"
var WatchedWalletsCollection = "watchedWallets"
var WatchedWalletsDoc = "watchedWallets"

var authFile = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

func init() {
	conf := &firebase.Config{ProjectID: projectID}

	ctx := context.Background()

	app, err := firebase.NewApp(ctx, conf, option.WithCredentialsFile(authFile))
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	firestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

func GetMintCollections(mint solana.PublicKey) (collections []*Collection, err error) {
	collectionsRaw, err := firestoreClient.Collection(CollectionCollection).Where("mints", "array-contains", mint.ToPointer().String()).Documents(context.Background()).GetAll()

	if err != nil {
		log.Fatalf("firestoreClient.Collection: %v", err)
	}

	for _, collection := range collectionsRaw {
		// decode collection
		var c Collection
		err = collection.DataTo(&c)
		if err != nil {
			log.Fatalf("collection.DataTo: %v", err)
		}
		collections = append(collections, &c)
	}

	// Filter duplicates
	var uniqueCollections = make(map[string]*Collection)
	for _, collection := range collections {
		uniqueCollections[collection.ID] = collection
	}

	collections = make([]*Collection, 0, len(uniqueCollections))
	for _, collection := range uniqueCollections {
		collections = append(collections, collection)
	}

	return collections, nil
}

func GetWatchedWallets() ([]*solana.PublicKey, error) {
	var watchedWallets []*solana.PublicKey
	watchedWalletsRef := firestoreClient.Collection(WatchedWalletsCollection).Doc(WatchedWalletsDoc)
	snap, err := watchedWalletsRef.Get(context.Background())
	if err != nil {
		return nil, err
	}
	var watchedWalletsObj WatchedWallets
	err = snap.DataTo(&watchedWalletsObj)
	if err != nil {
		return nil, err
	}
	for _, watchedWallet := range watchedWalletsObj.WatchedWallets {
		watchedWalletPtr, err := solana.PublicKeyFromBase58(watchedWallet)
		if err != nil {
			return nil, err
		}
		watchedWallets = append(watchedWallets, &watchedWalletPtr)
	}
	return watchedWallets, nil
}

func UpdateCollectionStats(collection string, stats map[string]*int64) error {
	collectionRef, err := firestoreClient.Collection(CollectionCollection).Doc(collection).Get(context.Background())
	if err != nil {
		return err
	}

	_, err = collectionRef.Ref.Set(context.Background(), stats, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func AddMintWithMetadata(mint solana.PublicKey, metadata Metadata) error {
	collectionRef := firestoreClient.Collection(MintsCollection)
	mintRef, _ := collectionRef.Doc(mint.String()).Get(context.Background())
	var firestoreCreators []FirestoreCreator
	for _, creator := range *metadata.Data.Creators {
		firestoreCreator := FirestoreCreator{
			Address:  creator.Address.String(),
			Verified: creator.Verified,
			Share:    creator.Share,
		}
		firestoreCreators = append(firestoreCreators, firestoreCreator)
	}
	if mintRef.Exists() {

	} else {
		mintData := map[string]interface{}{
			"isNFT":   true,
			"address": mint.String(),
			"metadata": map[string]interface{}{
				"key":             metadata.Key,
				"updateAuthority": metadata.UpdateAuthority.String(),
				"mint":            metadata.Mint.String(),
				"data": map[string]interface{}{
					"name":                 metadata.Data.Name,
					"symbol":               metadata.Data.Symbol,
					"uri":                  metadata.Data.Uri,
					"sellerFeeBasisPoints": metadata.Data.SellerFeeBasisPoints,
					"creators":             firestoreCreators,
				},
				"primarySaleHappened": metadata.PrimarySaleHappened,
				"isMutable":           metadata.IsMutable,
				"editionNonce":        metadata.EditionNonce,
				"collections":         []string{},
			},
		}

		batch := firestoreClient.Batch()
		batch.Set(mintRef.Ref, mintData, firestore.MergeAll)
		_, err := batch.Commit(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

func AddMintWithMetadataExtended(mint solana.PublicKey, metadata Metadata, extended interface{}) error {
	collectionRef := firestoreClient.Collection(MintsCollection)
	mintRef, _ := collectionRef.Doc(mint.String()).Get(context.Background())

	var firestoreCreators []FirestoreCreator
	for _, creator := range *metadata.Data.Creators {
		firestoreCreator := FirestoreCreator{
			Address:  creator.Address.String(),
			Verified: creator.Verified,
			Share:    creator.Share,
		}
		firestoreCreators = append(firestoreCreators, firestoreCreator)
	}
	// if mintRef.Exists() {

	// } else {
	mintData := map[string]interface{}{
		"isNFT":   true,
		"address": mint.String(),
		"metadata": map[string]interface{}{
			"key":             metadata.Key,
			"updateAuthority": metadata.UpdateAuthority.String(),
			"mint":            metadata.Mint.String(),
			"data": map[string]interface{}{
				"name":                 metadata.Data.Name,
				"symbol":               metadata.Data.Symbol,
				"uri":                  metadata.Data.Uri,
				"sellerFeeBasisPoints": metadata.Data.SellerFeeBasisPoints,
				"creators":             firestoreCreators,
			},
			"primarySaleHappened": metadata.PrimarySaleHappened,
			"isMutable":           metadata.IsMutable,
			"editionNonce":        metadata.EditionNonce,
			"extended":            extended,
		},
		"collections": []string{},
	}

	batch := firestoreClient.Batch()
	// batch.Create(mintRef.Ref, mintData)
	batch.Set(mintRef.Ref, mintData, firestore.MergeAll)

	_, err := batch.Commit(context.Background())
	if err != nil {
		return err
	}
	// }

	return nil
}

func AddOrder(order Order) error {
	orderCollectionRef := firestoreClient.Collection(OrdersCollection)
	// random id
	newOrderRef := orderCollectionRef.NewDoc()
	// m, err := json.Marshal(order)
	_, err := newOrderRef.Set(context.Background(), order)
	if err != nil {
		return err
	}

	return nil
}

func FindMagicEdenSellOrder(seller solana.PublicKey, assetMint solana.PublicKey, paymentMint solana.PublicKey, paymentSize uint64) (*string, error) {
	orderCollectionRef := firestoreClient.Collection(OrdersCollection)

	query := orderCollectionRef.Where("maker", "==", seller.String()).Where("assetMint", "==", assetMint.String()).Where("paymentMint", "==", paymentMint.String()).Where("paymentBaseSize", "==", int64(paymentSize)).Where("programID", "==", "M2mx93ekt1fmXSVkTrUL9xVFHkmME8HTUi5Cyc5aF7K").Where("side", "==", 1)
	iter := query.Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		return &doc.Ref.ID, nil
	}

	return nil, nil
}

func FindMagicEdenBuyOrder(buyer solana.PublicKey, assetMint solana.PublicKey, paymentMint solana.PublicKey, paymentSize uint64) (*string, error) {
	orderCollectionRef := firestoreClient.Collection(OrdersCollection)

	query := orderCollectionRef.Where("maker", "==", buyer.String()).Where("assetMint", "==", assetMint.String()).Where("paymentMint", "==", paymentMint.String()).Where("paymentBaseSize", "==", int64(paymentSize)).Where("programID", "==", "M2mx93ekt1fmXSVkTrUL9xVFHkmME8HTUi5Cyc5aF7K").Where("side", "==", 0)
	iter := query.Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		return &doc.Ref.ID, nil

	}

	return nil, nil
}

func SetOrderMatched(orderID string) error {
	orderCollectionRef := firestoreClient.Collection(OrdersCollection)
	orderRef := orderCollectionRef.Doc(orderID)
	_, err := orderRef.Update(context.Background(), []firestore.Update{{Path: "matched", Value: true}, {Path: "matchedDate", Value: int64(time.Now().Unix())}})
	if err != nil {
		return err
	}
	fmt.Printf("order %s matched\n", orderID)
	return nil
}

func SetOrderMatchedWith(orderID string, matchedWith string) error {
	orderCollectionRef := firestoreClient.Collection(OrdersCollection)
	orderRef := orderCollectionRef.Doc(orderID)
	_, err := orderRef.Update(context.Background(), []firestore.Update{
		{Path: "matched", Value: true},
		{Path: "matchedWith", Value: matchedWith},
		{Path: "matchedDate", Value: int64(time.Now().Unix())},
	})
	if err != nil {
		return err
	}
	fmt.Printf("order %s matched with %s\n", orderID, matchedWith)
	return nil
}

func CollectionExists(collectionKey solana.PublicKey) (bool, error) {
	collectionRef := firestoreClient.Collection(CollectionCollection)
	doc, _ := collectionRef.Doc(collectionKey.String()).Get(context.Background())
	if doc.Exists() {
		return true, nil
	}
	return false, nil
}

func AddMintToCollection(collectionKey solana.PublicKey, mintKey solana.PublicKey) error {
	collectionRef := firestoreClient.Collection(CollectionCollection)
	collectionDoc, err := collectionRef.Doc(collectionKey.String()).Get(context.Background())
	if err != nil {
		return errors.New("collection not found")
	}
	mintsRef := firestoreClient.Collection(MintsCollection)
	mintDoc, err := mintsRef.Doc(mintKey.String()).Get(context.Background())
	if err != nil {
		return errors.New("mint not found")
	}
	batch := firestoreClient.Batch()
	if collectionDoc.Exists() {
		if mintDoc.Exists() {
			batch.Update(collectionDoc.Ref, []firestore.Update{
				{Path: "mints", Value: firestore.ArrayUnion(mintKey.String())},
			})
			batch.Update(mintDoc.Ref, []firestore.Update{
				{Path: "collections", Value: firestore.ArrayUnion(collectionKey.String())},
			})
		} else {
			return errors.New("mint does not exist")
		}
	} else {
		return errors.New("collection does not exist")
	}
	_, err = batch.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("error adding mint %v to collection %v: %v", mintKey, collectionKey, err)
	}
	return nil
}

func CreateEmptyOnChainCollection(collectionKey solana.PublicKey, name string, description string, imageURL string) error {
	collectionRef := firestoreClient.Collection(CollectionCollection)
	collection := Collection{
		ID:                   collectionKey.String(),
		Type:                 "on-chain",
		User:                 "",
		Mints:                []string{},
		Name:                 name,
		Description:          description,
		Discord:              "",
		Website:              "",
		Twitter:              "",
		CreatedDate:          time.Now().String(),
		UpdatedDate:          time.Now().String(),
		CandyMachine:         "",
		OnChainCollectionKey: collectionKey.String(),
		Image:                imageURL,
	}

	docRef, err := collectionRef.Doc(collectionKey.String()).Get(context.Background())
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return err
	}
	if !docRef.Exists() {
		_, err := collectionRef.Doc(collectionKey.String()).Set(context.Background(), collection)
		if err != nil {
			return err
		}
	}
	fmt.Printf("created on-chain collection %s\n", collectionKey.String())
	return nil
}

func CreateEmptyCandyMachineCollection(candyMachineID solana.PublicKey, collectionName string) error {
	collectionRef := firestoreClient.Collection(CollectionCollection)
	collection := Collection{
		ID:                   candyMachineID.String(),
		Type:                 "candy-machine",
		User:                 "",
		Mints:                []string{},
		Name:                 collectionName,
		Description:          "",
		Discord:              "",
		Website:              "",
		Twitter:              "",
		CreatedDate:          time.Now().String(),
		UpdatedDate:          time.Now().String(),
		CandyMachine:         candyMachineID.String(),
		OnChainCollectionKey: "",
	}

	_, err := collectionRef.Doc(candyMachineID.String()).Set(context.Background(), collection)
	if err != nil {
		return err
	}
	fmt.Printf("created candy-machine collection %s\n", candyMachineID.String())
	return nil
}

func ReplaceCollectionMints(collectionKey solana.PublicKey, mints []solana.PublicKey) error {
	var stringifedPubkeys []string
	for _, mint := range mints {
		stringifedPubkeys = append(stringifedPubkeys, mint.String())
	}
	collectionRef := firestoreClient.Collection(CollectionCollection)
	_, err := collectionRef.Doc(collectionKey.String()).Update(context.Background(), []firestore.Update{{Path: "mints", Value: stringifedPubkeys}})
	if err != nil {
		return err
	}
	fmt.Printf("updated on-chain collection %s\n", collectionKey.String())
	return nil
}

func MintExists(mint solana.PublicKey) (bool, error) {
	mintRef := firestoreClient.Collection(MintsCollection).Doc(mint.String())
	doc, _ := mintRef.Get(context.Background())
	if doc.Exists() {
		return true, nil
	}
	return false, nil
}

func GetAllMints() ([]MintData, error) {
	mintsRef := firestoreClient.Collection(MintsCollection)
	mints := []MintData{}
	iter := mintsRef.Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var mint MintData
		err = doc.DataTo(&mint)
		if err != nil {
			return nil, err
		}
		mints = append(mints, mint)
	}
	return mints, nil
}

func WalletExists(wallet solana.PublicKey) (bool, error) {
	walletRef := firestoreClient.Collection(WalletsCollection).Doc(wallet.String())
	doc, _ := walletRef.Get(context.Background())
	if doc.Exists() {
		return true, nil
	}
	return false, nil
}

func CreateEmptyWalletDoc(wallet solana.PublicKey) error {
	walletRef := firestoreClient.Collection(WalletsCollection)
	walletObj := Wallet{
		Pubkey:             wallet.String(),
		OldestProcessedSig: "",
		LatestProcessedSig: "",
	}

	docRef, err := walletRef.Doc(walletObj.Pubkey).Get(context.Background())
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return err
	}
	if !docRef.Exists() {
		_, err := walletRef.Doc(wallet.String()).Set(context.Background(), walletObj)
		if err != nil {
			return err
		}
	}
	fmt.Printf("created wallet  %s\n", wallet.String())
	return nil
}

func GetWallet(wallet solana.PublicKey) (Wallet, error) {
	walletRef := firestoreClient.Collection(WalletsCollection).Doc(wallet.String())
	doc, err := walletRef.Get(context.Background())
	if err != nil {
		return Wallet{}, err
	}
	if !doc.Exists() {
		return Wallet{}, errors.New("wallet not found")
	}
	var walletObj Wallet
	err = doc.DataTo(&walletObj)
	if err != nil {
		return Wallet{}, err
	}
	return walletObj, nil
}

func SetWalletLatestProcessedSig(wallet solana.PublicKey, sig string) error {
	walletRef := firestoreClient.Collection(WalletsCollection).Doc(wallet.String())
	_, err := walletRef.Update(context.Background(), []firestore.Update{{Path: "latestProcessedSig", Value: sig}})
	if err != nil {
		return err
	}
	fmt.Printf("updated wallet latested processed sig %s\n", wallet.String())
	return nil
}

func SetWalletOldestProcessedSig(wallet solana.PublicKey, sig string) error {
	walletRef := firestoreClient.Collection(WalletsCollection).Doc(wallet.String())
	_, err := walletRef.Update(context.Background(), []firestore.Update{{Path: "oldestProcessedSig", Value: sig}})
	if err != nil {
		return err
	}
	fmt.Printf("updated wallet oldest processed sig %s\n", wallet.String())
	return nil
}

func SetWalletSyncedToSlot(wallet solana.PublicKey, slot int) error {
	walletRef := firestoreClient.Collection(WalletsCollection).Doc(wallet.String())
	_, err := walletRef.Update(context.Background(), []firestore.Update{{Path: "syncedToSlot", Value: slot}})
	if err != nil {
		return err
	}
	fmt.Printf("updated wallet synced to slot %s\n", wallet.String())
	return nil
}

func GetWalletSyncedToSlot(wallet solana.PublicKey) (int, error) {
	walletRef := firestoreClient.Collection(WalletsCollection).Doc(wallet.String())
	doc, err := walletRef.Get(context.Background())
	if err != nil {
		return 0, err
	}
	if !doc.Exists() {
		return 0, errors.New("wallet not found")
	}
	var walletObj Wallet
	err = doc.DataTo(&walletObj)
	if err != nil {
		return 0, err
	}
	return walletObj.SyncedToSlot, nil
}

func SetWalletOwnedTokens(wallet solana.PublicKey, ownedTokens []OwnedTokens) error {
	walletRef := firestoreClient.Collection(WalletsCollection).Doc(wallet.String())
	_, err := walletRef.Update(context.Background(), []firestore.Update{{Path: "ownedTokens", Value: ownedTokens}})
	if err != nil {
		return err
	}
	fmt.Printf("updated wallet owned tokens %s\n", wallet.String())
	return nil
}

func SetSigProcessedForWallet(wallet solana.PublicKey, sig solana.Signature) error {
	walletRef := firestoreClient.Collection(WalletsSigsCollection).Doc(wallet.String())
	// _, err := walletRef.Set(context.Background(), map[string]interface{}{"sigsProcessed": sig.String()}, firestore.MergeAll)
	firstData := map[string]interface{}{
		"address": wallet.String(),
	}
	batch := firestoreClient.Batch()
	batch.Set(walletRef, firstData, firestore.MergeAll)
	batch.Update(walletRef, []firestore.Update{
		{Path: "processedSigs", Value: firestore.ArrayUnion(sig.String())},
	})
	_, err := batch.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("error updating wallet sig %s: %s", wallet.String(), err.Error())
	}
	fmt.Printf("updated wallet sig processed %s\n", wallet.String())
	return nil
}

func SetSigProcessed(sig solana.Signature) error {
	sigRef := firestoreClient.Collection(ProcessedSigsCollection).Doc(sig.String())
	time := time.Now().Unix()
	firstData := map[string]interface{}{
		"signature": sig.String(),
		"time":      time,
	}
	_, err := sigRef.Set(context.Background(), firstData, firestore.MergeAll)
	if err != nil {
		return fmt.Errorf("error updating sig %s: %s", sig.String(), err.Error())
	}
	return nil
}

func GetSigIsProcessed(sig solana.Signature) (bool, error) {
	sigRef := firestoreClient.Collection(ProcessedSigsCollection).Doc(sig.String())
	doc, err := sigRef.Get(context.Background())
	if err != nil {
		return false, err
	}
	if !doc.Exists() {
		return false, nil
	}
	return true, nil
}

func GetSigIsProcessedForWallet(wallet solana.PublicKey, sig solana.Signature) (bool, error) {
	walletRef := firestoreClient.Collection(WalletsSigsCollection).Doc(wallet.String())
	doc, err := walletRef.Get(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		}
		return false, err
	}
	if !doc.Exists() {
		return false, errors.New("wallet not found")
	}
	var walletObj WalletsProcessed
	err = doc.DataTo(&walletObj)
	if err != nil {
		return false, err
	}
	for _, sigProcessed := range walletObj.ProcessedSigs {
		if sigProcessed == sig.String() {
			return true, nil
		}
	}
	return false, nil
}
