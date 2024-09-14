package firebase

import (
	"context"
	"fmt"
	"log"
	"os"

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
		credentialsFilePath = "third_party/gcloud/firebase/service_account.json"
	}

	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(credentialsFilePath))
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}

	appInstance = app

	InitAuth(*app)
	firebase_storage.Init(*app)
}

func GetApp() *firebase.App{
    return appInstance
}