package api

import (
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api/handlers"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
)

// SetupRouter initializes the HTTP router
func SetupRouter(prescriptionModel *models.PrescriptionModel, medicationModel *models.MedicationModel) *http.ServeMux {
	mux := http.NewServeMux()

	// Prescription Handlers
	prescriptionHandler := handlers.NewPrescriptionHandler(prescriptionModel)
	mux.HandleFunc("/prescriptions", prescriptionHandler.GetAllPrescriptionsHandler) // To GET all prescriptions
	mux.HandleFunc("/prescriptions", prescriptionHandler.AddPrescription)            // To POST a new prescription

	// Medication Handlers
	medicationHandler := handlers.NewMedicationHandler(medicationModel)
	mux.HandleFunc("/medications", medicationHandler.GetAllMedications) // To GET all medications
	mux.HandleFunc("/medications", medicationHandler.AddMedication)     // To POST a new medication

	// Upload Route
	uploadHandler := handlers.NewUploadHandler()                // Assuming an UploadHandler exists
	mux.HandleFunc("/upload", uploadHandler.UploadImageHandler) // POST image to the server

	// Default route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "API is running. Use /prescriptions to fetch data, /medications to fetch medications.", http.StatusNotFound)
	})

	return mux
}
