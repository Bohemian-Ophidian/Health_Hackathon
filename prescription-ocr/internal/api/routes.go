package api

import (
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api/handlers"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
)

// SetupRouter initializes the HTTP router
func SetupRouter(prescriptionModel *models.PrescriptionModel) *http.ServeMux {
	mux := http.NewServeMux()

	prescriptionHandler := handlers.NewPrescriptionHandler(prescriptionModel)

	mux.HandleFunc("/prescriptions", prescriptionHandler.GetAllPrescriptionsHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "API is running. Use /prescriptions to fetch data.", http.StatusNotFound)
	})

	return mux
}
