package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MedicationHandler manages medication routes
type MedicationHandler struct {
	Model *models.MedicationModel
}

// NewMedicationHandler initializes the handler
func NewMedicationHandler(model *models.MedicationModel) *MedicationHandler {
	return &MedicationHandler{Model: model}
}

// GetMedication retrieves a medication by ID
func (h *MedicationHandler) GetMedication(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	medication, err := h.Model.GetMedication(ctx, id)
	if err != nil {
		http.Error(w, "Medication not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medication)
}

// CreateMedication adds a new medication
func (h *MedicationHandler) CreateMedication(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var medication models.Medication
	if err := json.NewDecoder(r.Body).Decode(&medication); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	medication.ID = primitive.NewObjectID() // Correct way to assign ObjectID
	medication.CreatedAt = time.Now()
	medication.UpdatedAt = time.Now()

	if err := h.Model.AddMedication(ctx, &medication); err != nil {
		http.Error(w, "Failed to create medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(medication)
}

// UpdateMedication updates an existing medication
func (h *MedicationHandler) UpdateMedication(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	var medication models.Medication
	if err := json.NewDecoder(r.Body).Decode(&medication); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Model.UpdateMedication(ctx, id, &medication); err != nil {
		http.Error(w, "Failed to update medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Medication updated successfully"})
}

// DeleteMedication deletes a medication
func (h *MedicationHandler) DeleteMedication(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	if err := h.Model.DeleteMedication(ctx, id); err != nil {
		http.Error(w, "Failed to delete medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Medication deleted successfully"})
}