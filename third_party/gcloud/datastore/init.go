package datastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

type DatastoreClient struct {
	*datastore.Client
}

var datastoreClient DatastoreClient

// Init initializes the Datastore client.
func Init(projectID string) {
	ctx := context.Background()

	// Create a new Datastore client.
	client, err := datastore.NewClient(ctx, projectID, option.WithCredentialsFile("./third_party/gcloud/firebase/service_account_dev.json"))
	if err != nil {
		log.Fatalf("Failed to create Datastore client: %v", err)
	}

	datastoreClient = DatastoreClient{client}
}

// GetDatastoreClient returns the initialized Datastore client.
func GetDatastoreClient() DatastoreClient {
	return datastoreClient
}
