package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
)

// MedicationHandler handles medication API
type MedicationHandler struct {
	Model *models.MedicationModel
}

// NewMedicationHandler initializes the handler
func NewMedicationHandler(model *models.MedicationModel) *MedicationHandler {
	return &MedicationHandler{Model: model}
}

// HandleMedications handles different HTTP methods for medications
func (h *MedicationHandler) HandleMedications(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.AddMedication(w, r)
	case http.MethodPut:
		h.UpdateMedication(w, r)
	case http.MethodDelete:
		h.DeleteMedication(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// AddMedication inserts a new medication into the database
func (h *MedicationHandler) AddMedication(w http.ResponseWriter, r *http.Request) {
	var medication models.Medication
	err := json.NewDecoder(r.Body).Decode(&medication)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = h.Model.AddMedication(ctx, &medication)
	if err != nil {
		http.Error(w, "Failed to add medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(medication)
}

// UpdateMedication updates an existing medication
func (h *MedicationHandler) UpdateMedication(w http.ResponseWriter, r *http.Request) {
	var medication models.Medication
	err := json.NewDecoder(r.Body).Decode(&medication)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = h.Model.UpdateMedication(ctx, medication.ID.Hex(), &medication)
	if err != nil {
		http.Error(w, "Failed to update medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(medication)
}

// DeleteMedication removes a medication from the database
func (h *MedicationHandler) DeleteMedication(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ID string `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = h.Model.DeleteMedication(ctx, request.ID)
	if err != nil {
		http.Error(w, "Failed to delete medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Medication deleted successfully"})
}
