package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
)

// UploadsHandler handles prescription uploads
type UploadsHandler struct {
	PrescriptionModel *models.PrescriptionModel
}

// NewUploadsHandler initializes the uploads handler
func NewUploadsHandler(prescriptionModel *models.PrescriptionModel) *UploadsHandler {
	return &UploadsHandler{PrescriptionModel: prescriptionModel}
}

// HandleUploads processes prescription uploads
func (h *UploadsHandler) HandleUploads(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var prescription models.Prescription
	err := json.NewDecoder(r.Body).Decode(&prescription)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = h.PrescriptionModel.AddPrescription(ctx, &prescription)
	if err != nil {
		http.Error(w, "Failed to save prescription", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prescription)
}
