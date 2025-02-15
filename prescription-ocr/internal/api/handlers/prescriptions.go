package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
)

// PrescriptionHandler is the struct for handling prescriptions
type PrescriptionHandler struct {
	Model *models.PrescriptionModel
}

// NewPrescriptionHandler creates a new PrescriptionHandler
func NewPrescriptionHandler(model *models.PrescriptionModel) *PrescriptionHandler {
	return &PrescriptionHandler{Model: model}
}

// GetAllPrescriptionsHandler retrieves all prescriptions
func (h *PrescriptionHandler) GetAllPrescriptionsHandler(w http.ResponseWriter, r *http.Request) {
	prescriptions, err := h.Model.GetAllPrescriptions(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve prescriptions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prescriptions)
}

// AddPrescription handles adding a prescription
func (h *PrescriptionHandler) AddPrescription(w http.ResponseWriter, r *http.Request) {
	var prescription models.Prescription
	if err := json.NewDecoder(r.Body).Decode(&prescription); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.Model.AddPrescription(r.Context(), &prescription)
	if err != nil {
		http.Error(w, "Failed to add prescription", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Prescription added successfully"))
}
