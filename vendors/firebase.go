package vendors

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/SamPariatIL/weather-wrapper/config"
	"google.golang.org/api/option"
	"log"
)

var firebaseApp *firebase.App
var firebaseAuth *auth.Client

func InitFirebaseAdmin() {
	conf := config.GetConfig()
	ctx := context.Background()

	credentialsJson, err := json.Marshal(conf.FirebaseConfig)
	if err != nil {
		log.Fatalf("Failed to marshal Firebase credentials: %v", err)
	}

	adminOpt := option.WithCredentialsJSON(credentialsJson)

	firebaseApp, err = firebase.NewApp(ctx, nil, adminOpt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase admin sdk: %v", err)
	}

	firebaseAuth, err = firebaseApp.Auth(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase auth client: %v", err)
	}

	log.Println("Initialized Firebase!")
}

func GetFirebaseApp() *firebase.App {
	return firebaseApp
}

func GetFirebaseAuth() *auth.Client {
	return firebaseAuth
}
