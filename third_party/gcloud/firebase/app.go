package firebase

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bakeable/bkry/third_party/gcloud/firebase/firestore"
	firebase_storage "github.com/bakeable/bkry/third_party/gcloud/firebase/storage"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var appInstance *firebase.App

func Init() {
	var credentialsFilePath string
	environment := os.Getenv("GO_ENV")
	fmt.Println("Environment: " + environment)
	
	if environment == "development" {
		credentialsFilePath = "third_party/gcloud/firebase/service_account_dev.json"
	} else {
		credentialsFilePath = "third_party/gcloud/firebase/service_account_staging.json"
	}

	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: "brain-bash-staging",
	}, option.WithCredentialsFile(credentialsFilePath))
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}
	fmt.Println("Firebase app initialized", credentialsFilePath)

	appInstance = app

	InitAuth(*app)
	firestore.Init(*app)
	firebase_storage.Init(*app)
}

func GetApp() *firebase.App{
    return appInstance
}