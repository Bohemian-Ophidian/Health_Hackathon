package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/config"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/ollama" // Import LLaMA client
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠ No .env file found, using system environment variables")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// MongoDB connection
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}
	err = client.Connect(nil)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(nil)

	// Verify the connection to healthDB
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME environment variable is not set")
	}
	database := client.Database(dbName)

	// Initialize models
	prescriptionModel := models.NewPrescriptionModel(database)
	medicationModel := models.NewMedicationModel(database)

	// Initialize LLaMA client
	llamaClient := llama.NewClient(cfg.LLaMA.APIURL, os.Getenv("LLAMA_API_KEY"))

	// Set up the router
	router := api.SetupRouter(prescriptionModel, medicationModel, database, llamaClient)

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("🚀 Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
