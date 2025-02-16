package ocr_test

import (
	"testing"

<<<<<<< HEAD
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/ocr" // Replace `your_project` with your actual module name
=======
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/ocr"
>>>>>>> 6b5b2b5ce4434bd8b80f655908dba12693114995
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
