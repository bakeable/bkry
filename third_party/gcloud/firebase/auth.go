package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var authClient *auth.Client

func InitAuth(app firebase.App) {
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
	}

	authClient = auth
}

func GetAuthClient() *auth.Client{
    return authClient
}

func GetPasswordResetLink(email string) string {
	link, err := GetAuthClient().PasswordResetLinkWithSettings(context.Background(), email, nil)
	if err != nil {
		panic(err)
	}

	return link
}