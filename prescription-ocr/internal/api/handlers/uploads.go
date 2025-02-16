package handlers

import (
	"net/http"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/common" // Import the common interface
)

type UploadHandler struct {
	LLaMAClient common.LLaMAClient
}

func NewUploadHandler(llamaClient common.LLaMAClient) *UploadHandler {
	return &UploadHandler{LLaMAClient: llamaClient}
}

func (h *UploadHandler) UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	// Your implementation
}
