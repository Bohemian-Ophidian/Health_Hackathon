package ocr

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/llama"
)

// ProcessImage extracts text using OCR and refines it with LLaMA
func ProcessImage(imagePath string) (string, error) {
	// Check if file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", imagePath)
	}

	log.Println("🛠️ Running Tesseract on:", imagePath)

	// Run Tesseract OCR
	cmd := exec.Command("tesseract", imagePath, "stdout", "--psm", "6", "--oem", "3")
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Tesseract Error: %s\n", string(output))
		return "", fmt.Errorf("failed to process image: %s", string(output))
	}

	extractedText := string(output)
	log.Println("✅ Extracted Text:", extractedText)

	// Initialize LLaMA client
	llamaClient := llama.NewClient("http://localhost:8080") // Use your actual LLaMA API URL

	// Send extracted text to LLaMA for further processing
	analysis, err := llamaClient.AnalyzeMedication(context.Background(), extractedText)
	if err != nil {
		log.Printf("LLaMA Processing Failed: %v", err)
		return extractedText, nil // Return raw extracted text if LLaMA fails
	}

	log.Println("✅ Analyzed Medication:", analysis)
	// You can use the analysis data as needed or return it
	return fmt.Sprintf("Analysis Result: %v", analysis), nil
}
