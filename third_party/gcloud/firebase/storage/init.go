package firebase_storage

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
)

type FirebaseStorageClient struct {
	*storage.Client
}

var firebaseStorageClient FirebaseStorageClient

func Init(app firebase.App) {
	ctx := context.Background()

	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firebase Storage client: %v", err)
	}

	firebaseStorageClient = FirebaseStorageClient{client}
}

func GetFirebaseStorageClient() FirebaseStorageClient {
	return firebaseStorageClient
}