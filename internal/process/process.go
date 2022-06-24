package process

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/lukejmann/carmel-backend/internal/constants"
	"google.golang.org/api/option"
)

var rpcClient *rpc.Client
var bqClient *bigquery.Client
var firebaseClient *firestore.Client
var httpClient = http.Client{}

var projectID = "nice-1d36b"
var datasetID = "niceDevnet"

var authFile = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

func init() {
	conf := &firebase.Config{ProjectID: projectID}

	ctx := context.Background()

	app, err := firebase.NewApp(ctx, conf, option.WithCredentialsFile(authFile))
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	firebaseClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}

	endpoint := constants.RPCEndPoint
	rpcClient = rpc.New(endpoint)
	bqClient, _ = bigquery.NewClient(context.Background(), projectID)
	defer bqClient.Close()

	httpClient = http.Client{
		Timeout: time.Second * 30, // Timeout after 2 seconds
	}
}
