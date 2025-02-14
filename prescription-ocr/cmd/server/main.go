package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/utils"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠ No .env file found, using system environment variables")
	}

	// Required environment variables
	requiredEnvs := []string{"MONGO_URI", "DB_NAME", "SERVER_PORT"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			log.Fatalf("❌ Missing required environment variable: %s", env)
		}
	}

	// Connect to MongoDB
	db, err := utils.ConnectMongoDB(os.Getenv("MONGO_URI"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}
	log.Println("✅ Connected to MongoDB successfully")

	// Initialize models
	// medicationModel := models.NewMedicationModel(db)
	PrescriptionModel := models.NewPrescriptionModel(db)

	// Setup router
	router := api.SetupRouter(PrescriptionModel)

	// Start server
	port := os.Getenv("SERVER_PORT")
	log.Printf("🚀 Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
