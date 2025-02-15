package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
)

// MedicationHandler is the struct for the medication handler
type MedicationHandler struct {
	Model *models.MedicationModel
}

// NewMedicationHandler creates a new MedicationHandler
func NewMedicationHandler(model *models.MedicationModel) *MedicationHandler {
	return &MedicationHandler{Model: model}
}

// GetAllMedications handles GET requests to retrieve all medications
func (h *MedicationHandler) GetAllMedications(w http.ResponseWriter, r *http.Request) {
	medications, err := h.Model.GetAllMedications(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve medications", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medications)
}

// AddMedication handles POST requests to add a new medication
func (h *MedicationHandler) AddMedication(w http.ResponseWriter, r *http.Request) {
	var medication models.Medication
	if err := json.NewDecoder(r.Body).Decode(&medication); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.Model.AddMedication(r.Context(), &medication)
	if err != nil {
		http.Error(w, "Failed to add medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Medication added successfully"))
}
