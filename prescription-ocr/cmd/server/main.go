import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/config"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/llama" // Added the LLaMA import here
)

func main() {
	// Load .env file (if present)
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
	if database == nil {
		log.Fatalf("Failed to access database: %s", dbName)
	} else {
		fmt.Printf("Successfully connected to MongoDB database: %s\n", dbName)
	}

	// Initialize models
	prescriptionModel := models.NewPrescriptionModel(database)
	medicationModel := models.NewMedicationModel(database)

	// Set up the router
	router := api.SetupRouter(prescriptionModel, medicationModel, database) // Pass database here

	// You can now access LLaMA API information through the cfg variable
	// For example, log the LLaMA API URL
	fmt.Println("LLaMA API URL:", cfg.LLaMA.APIURL)

	// Call the LLaMA API for any necessary processing here

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
