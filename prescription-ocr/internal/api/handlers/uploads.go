package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/ocr"
)

// UploadHandler handles image uploads
type UploadHandler struct{}

// NewUploadHandler creates a new UploadHandler
func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

// UploadImageHandler handles the upload of an image
func (h *UploadHandler) UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	// Limit file size to 10 MB
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

	// Save the file to the local filesystem
	dst, err := os.Create("./uploads/uploaded_image.png")
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

	// Process the image (e.g., OCR, etc.)
	extractedText, err := ocr.ProcessImage("./uploads/uploaded_image.png")
	if err != nil {
		http.Error(w, "Failed to process image", http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("File uploaded successfully. Extracted text: %s", extractedText)))
}
