package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
)

// PrescriptionHandler handles prescription API operations
type PrescriptionHandler struct {
	Model *models.PrescriptionModel
}

// NewPrescriptionHandler initializes the handler
func NewPrescriptionHandler(model *models.PrescriptionModel) *PrescriptionHandler {
	return &PrescriptionHandler{Model: model}
}

// GetAllPrescriptionsHandler retrieves all prescriptions from the database
func (h *PrescriptionHandler) GetAllPrescriptionsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	prescriptions, err := h.Model.GetAllPrescriptions(ctx)
	if err != nil {
		http.Error(w, "Failed to retrieve prescriptions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prescriptions)
}
