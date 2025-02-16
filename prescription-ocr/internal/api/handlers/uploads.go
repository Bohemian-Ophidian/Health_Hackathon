package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/ollama"
)

// UploadHandler handles file uploads and OCR
type UploadHandler struct {
	OllamaClient *ollama.Client
}

// NewUploadHandler creates a new upload handler with Ollama client
func NewUploadHandler(ollamaClient *ollama.Client) *UploadHandler {
	return &UploadHandler{
		OllamaClient: ollamaClient,
	}
}

// UploadImageHandler handles the file upload and triggers OCR
func (h *UploadHandler) UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	// Your implementation to handle file upload and Tesseract OCR...
	// Extract text from the image (using Tesseract or any OCR library you prefer)

	// For simplicity, let's say the text extracted is "Aspirin"
	extractedText := "Aspirin"

	// Use Ollama to analyze the medication
	analysis, err := h.OllamaClient.AnalyzeMedication(r.Context(), extractedText)
	if err != nil {
		http.Error(w, "Failed to analyze medication with Ollama API", http.StatusInternalServerError)
		return
	}

	// Respond with analysis result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(analysis)
}
