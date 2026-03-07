package config

import (
	"log"
	"os"

	brevo "github.com/getbrevo/brevo-go/lib"
)

func NewBrevoClient() *brevo.APIClient {
	apiKey := os.Getenv("BREVO_API_KEY")
	if apiKey == "" {
		log.Fatal("FATAL: BREVO_API_KEY environment variable is missing")
	}

	cfg := brevo.NewConfiguration()
	cfg.AddDefaultHeader("api-key", apiKey)
	
	log.Println("Brevo SDK successfully initialized.")
	return brevo.NewAPIClient(cfg)
}