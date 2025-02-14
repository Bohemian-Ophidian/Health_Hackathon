package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/database"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/llama"
)

var mongoDB *database.MongoDB

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠ No .env file found, using system environment variables")
	}

	requiredEnvs := []string{"MONGO_URI", "MONGO_DB_NAME", "LLAMA_API_URL"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			log.Fatalf("❌ Missing required environment variable: %s", env)
		}
	}

	// Connect to MongoDB
	var err error
	mongoDB, err = database.ConnectMongoDB(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB_NAME"))
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Client.Disconnect(context.TODO())

	llamaClient := llama.NewClient(os.Getenv("LLAMA_API_URL"))

	// Initialize models
	medicationModel := &models.MedicationModel{DB: mongoDB.DB}

	// Setup API router
	router := api.SetupRouter(medicationModel, llamaClient)

	// Start HTTP server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}