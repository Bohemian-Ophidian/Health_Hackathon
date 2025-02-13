package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/database"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/llama"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠ No .env file found, using system environment variables")
	}

	requiredEnvs := []string{"MONGO_URI", "MONGO_DB_NAME", "SERVER_PORT", "LLAMA_API_URL"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			log.Fatalf("❌ Missing required environment variable: %s", env)
		}
	}

	mongoDB, err := database.ConnectMongoDB()
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	medicationModel := models.NewMedicationModel(mongoDB)
	llamaClient := llama.NewClient(os.Getenv("LLAMA_API_URL"))

	router := api.SetupRouter(medicationModel, llamaClient)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}