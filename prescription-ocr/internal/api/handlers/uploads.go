package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/models"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/llama"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/ocr"
)

// UploadHandler handles the image uploads
type UploadHandler struct {
	LLaMAClient     *llama.Client
	MedicationModel *models.MedicationModel
}

// NewUploadHandler creates a new UploadHandler
func NewUploadHandler(llamaClient *llama.Client, medicationModel *models.MedicationModel) *UploadHandler {
	return &UploadHandler{
		LLaMAClient:     llamaClient,
		MedicationModel: medicationModel,
	}
}

// UploadImageHandler handles image uploads, processes them, and returns the response
func (h *UploadHandler) UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	// Limit file size to 10MB
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

	// Save the uploaded file temporarily
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

	// Send extracted text to Ollama API for medication analysis
	analysisResult, err := h.LLaMAClient.AnalyzeMedication(context.Background(), extractedText)
	if err != nil {
		http.Error(w, "Failed to analyze medication with Ollama API", http.StatusInternalServerError)
		return
	}

	// Search the database for the medication details
	medication, err := h.MedicationModel.GetMedicationDetails(r.Context(), analysisResult["medication_name"].(string))
	if err != nil {
		http.Error(w, "Failed to retrieve medication details", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	response := map[string]interface{}{
		"message":         "File uploaded and processed successfully",
		"extracted_text":  extractedText,
		"medication_info": medication,
	}

	// Return the response to the user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
