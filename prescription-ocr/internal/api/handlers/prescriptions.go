package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/ocr"
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/ollama"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	// Handle file upload and save it
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the uploaded file
	outFile, err := os.Create("uploads/prescription.jpg")
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = outFile.ReadFrom(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Send a response that the file was uploaded
	w.Write([]byte("File uploaded successfully"))
}

func ProcessPrescription(w http.ResponseWriter, r *http.Request) {
	var input struct {
		PrePrompt string `json:"pre_prompt"`
	}

	// Decode the input JSON containing the pre-prompt
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Extract text from the prescription image
	text := ocr.ExtractTextFromImage("uploads/prescription.jpg")

	// Combine pre-prompt with extracted text for Ollama processing
	combinedText := input.PrePrompt + " " + text

	// Process the combined text using Ollama
	medicationInfo := ollama.ProcessPrescription(combinedText)

	// Respond with the medication information
	response := map[string]string{"medicationInfo": medicationInfo}
	json.NewEncoder(w).Encode(response)
}
