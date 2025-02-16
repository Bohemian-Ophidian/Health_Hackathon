package config

import (
	"log"
	"os"
)

func GetOllamaAPIEndpoint() string {
	endpoint := os.Getenv("OLLAMA_API_ENDPOINT")
	if endpoint == "" {
		log.Fatal("OLLAMA_API_ENDPOINT not set in environment variables")
	}
	return endpoint
}
