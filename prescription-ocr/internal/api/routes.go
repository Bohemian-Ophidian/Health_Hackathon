package api

import (
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api/handlers"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/llama"
)

func SetupRouter(model *models.MedicationModel, llamaClient *llama.Client) *http.ServeMux {
	mux := http.NewServeMux()

	medicationHandler := handlers.NewMedicationHandler(model, llamaClient)

	mux.HandleFunc("/medications", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			medicationHandler.CreateMedication(w, r)
		case http.MethodGet:
			medicationHandler.GetMedication(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "API is running. Use /medications", http.StatusNotFound)
	})

	return mux
}