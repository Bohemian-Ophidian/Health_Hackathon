package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/database"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠ No .env file found, using system environment variables")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	mongoDB, err := database.ConnectMongoDB(mongoURI, dbName)
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Client.Disconnect(nil)

	// Initialize models
	prescriptionModel := models.NewPrescriptionModel(mongoDB.DB)

	// Setup API routes
	router := api.SetupRouter(prescriptionModel)

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}