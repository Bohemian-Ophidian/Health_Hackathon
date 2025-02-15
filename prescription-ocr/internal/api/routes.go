package api

import (
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api/handlers"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
)

// SetupRouter initializes the HTTP router
func SetupRouter(prescriptionModel *models.PrescriptionModel, medicationModel *models.MedicationModel) *http.ServeMux {
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

	// Upload Route
	uploadHandler := handlers.NewUploadHandler()                // Assuming an UploadHandler exists
	mux.HandleFunc("/upload", uploadHandler.UploadImageHandler) // POST image to the server

	// Default route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "API is running. Use /prescriptions to fetch data, /medications to fetch medications.", http.StatusNotFound)
	})

	return mux
}
