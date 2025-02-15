package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/llama"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/ocr"
	"go.mongodb.org/mongo-driver/mongo"
)

type UploadHandler struct {
	database *mongo.Database
}

func NewUploadHandler(database *mongo.Database) *UploadHandler {
	return &UploadHandler{database: database}
}

func (h *UploadHandler) UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	const MAX_UPLOAD_SIZE = 10 << 20
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	// Parse the multipart form data
	err := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the uploaded file
	dstPath := "./uploads/uploaded_image.png"
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Extract text using Tesseract
	extractedText, err := ocr.ProcessImage(dstPath)
	if err != nil {
		http.Error(w, "Failed to process image", http.StatusInternalServerError)
		return
	}

	// Initialize LLaMA client and analyze the extracted text for medication names
	llamaClient := llama.NewClient("http://localhost:8080")
	medicationNames, err := llamaClient.AnalyzeMedicationNames(extractedText)
	if err != nil {
		http.Error(w, "Failed to analyze medication names", http.StatusInternalServerError)
		return
	}

	// Fetch medication details from the database
	var medicationsDetails []models.MedicationDetails
	for _, medicationName := range medicationNames {
		medication, err := models.NewMedicationModel(h.database).GetMedicationDetails(r.Context(), medicationName)
		if err != nil {
			http.Error(w, "Failed to retrieve medication details", http.StatusInternalServerError)
			return
		}
		medicationsDetails = append(medicationsDetails, *medication)
	}

	// Prepare response with extracted text and medication details
	response := map[string]interface{}{
		"message":         "File uploaded successfully",
		"extracted_text":  extractedText,
		"medication_info": medicationsDetails,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
