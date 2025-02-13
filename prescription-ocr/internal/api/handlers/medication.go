package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MedicationHandler struct {
	Model *models.MedicationModel
}

func NewMedicationHandler(model *models.MedicationModel) *MedicationHandler {
	return &MedicationHandler{Model: model}
}

func (h *MedicationHandler) GetMedication(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	medication, err := h.Model.GetMedication(id)
	if err != nil {
		http.Error(w, "Medication not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(medication)
}

func (h *MedicationHandler) CreateMedication(w http.ResponseWriter, r *http.Request) {
	var medication models.Medication
	if err := json.NewDecoder(r.Body).Decode(&medication); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	medication.ID = primitive.NewObjectID().Hex() // Generate MongoDB Object ID
	if err := h.Model.AddMedication(&medication); err != nil {
		http.Error(w, "Failed to create medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(medication)
}