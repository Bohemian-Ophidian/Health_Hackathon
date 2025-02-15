package api

import (
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api/handlers"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRouter initializes the HTTP router
func SetupRouter(prescriptionModel *models.PrescriptionModel, medicationModel *models.MedicationModel, database *mongo.Database) *http.ServeMux {
	mux := http.NewServeMux()

	// Prescription Handlers (Handle both GET and POST)
	prescriptionHandler := handlers.NewPrescriptionHandler(prescriptionModel)
	mux.HandleFunc("/prescriptions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			prescriptionHandler.GetAllPrescriptionsHandler(w, r)
		case "POST":
			prescriptionHandler.AddPrescription(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Medication Handlers (Handle both GET and POST)
	medicationHandler := handlers.NewMedicationHandler(medicationModel)
	mux.HandleFunc("/medications", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			medicationHandler.GetAllMedications(w, r)
		case "POST":
			medicationHandler.AddMedication(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Upload Route (Pass the database to the UploadHandler)
	uploadHandler := handlers.NewUploadHandler(database)        // Pass database here
	mux.HandleFunc("/upload", uploadHandler.UploadImageHandler) // POST image to the server

	// Default route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "API is running. Use /prescriptions to fetch data, /medications to fetch medications.", http.StatusNotFound)
	})

	return mux
}
