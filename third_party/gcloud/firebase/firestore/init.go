package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type FirestoreClient struct {
	*firestore.Client
}

var firestoreClient FirestoreClient

func Init(app firebase.App) {
	ctx := context.Background()

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	firestoreClient = FirestoreClient{client}
}

func GetFirestoreClient() FirestoreClient {
	return firestoreClient
}