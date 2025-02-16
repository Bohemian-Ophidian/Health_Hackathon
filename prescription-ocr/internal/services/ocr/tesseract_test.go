package ocr

import (
	"testing"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/ocr"
)

func TestProcessImage(t *testing.T) {
	imagePath := "E://Codes//Projects//Health_Hackathon//prescription-ocr//cmd//server//uploads//sample_text.png"

	text, err := ocr.ProcessImage(imagePath)
	if err != nil {
		t.Fatalf("ProcessImage failed: %v", err)
	}

	if text == "" {
		t.Fatalf("Expected extracted text but got empty string")
	}
}
