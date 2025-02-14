package api

import (
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api/handlers"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
)

// SetupRouter initializes routes
func SetupRouter(prescriptionModel *models.PrescriptionModel) *http.ServeMux {
	mux := http.NewServeMux()

	// Initialize handlers
	prescriptionHandler := handlers.NewPrescriptionHandler(prescriptionModel)

	mux.HandleFunc("/upload", prescriptionHandler.UploadImageHandler)
	mux.HandleFunc("/prescriptions", prescriptionHandler.GetPrescriptionsHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "API is running", http.StatusNotFound)
	})

	return mux
}