package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
)

// PrescriptionHandler struct
type PrescriptionHandler struct {
	Model *models.PrescriptionModel
}

// NewPrescriptionHandler initializes handler
func NewPrescriptionHandler(model *models.PrescriptionModel) *PrescriptionHandler {
	return &PrescriptionHandler{Model: model}
}

// UploadImageHandler handles image uploads and OCR processing
func (h *PrescriptionHandler) UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		http.Error(w, "❌ Error parsing form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "❌ Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Process image using OCR (mocked function)
	extractedText := "Mock extracted text from image"

	// Save to MongoDB
	prescription := &models.Prescription{
		OriginalImage: "sample_image.png", // In real case, store actual file name/path
		ExtractedText: extractedText,
	}

	err = h.Model.SavePrescription(context.Background(), prescription)
	if err != nil {
		http.Error(w, "❌ Error saving prescription", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(prescription)
}

// GetPrescriptionsHandler fetches prescriptions
func (h *PrescriptionHandler) GetPrescriptionsHandler(w http.ResponseWriter, r *http.Request) {
	prescriptions, err := h.Model.GetPrescriptions(context.Background())
	if err != nil {
		http.Error(w, "❌ Failed to retrieve prescriptions", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(prescriptions)
}