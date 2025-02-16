package ocr

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
<<<<<<< HEAD
	"path/filepath"
)

// ProcessImage extracts text using OCR
func ProcessImage(imagePath string) (string, error) {
	// Get the absolute path of the image file
	absPath, err := filepath.Abs(imagePath)
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
		return "", fmt.Errorf("failed to get absolute path: %v", err)
	}
	log.Println("🛠️ Running Tesseract on:", absPath)

	// Check if the image file exists before processing
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", absPath)
	}

	// Command to run Tesseract, using the absolute path of the image
	cmd := exec.Command("tesseract", absPath, "stdout", "--psm", "6", "--oem", "3")

	// Run the command and capture both standard output and error
=======

	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/services/common" // Import the common interface
)

func ProcessImage(imagePath string, llamaClient common.LLaMAClient) (string, error) {
	// Check if file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", imagePath)
	}

	log.Println("🛠️ Running Tesseract on:", imagePath)

	// Run Tesseract OCR
	cmd := exec.Command("tesseract", imagePath, "stdout", "--psm", "6", "--oem", "3")
>>>>>>> 6b5b2b5ce4434bd8b80f655908dba12693114995
	output, err := cmd.CombinedOutput()

	// Error handling
	if err != nil {
		log.Printf("Tesseract Error: %s\n", string(output))
		return "", fmt.Errorf("failed to process image: %s", string(output))
	}

<<<<<<< HEAD
	// Log the extracted text (optional)
	log.Println("✅ Extracted Text: ", string(output))

	// Return the extracted text
	return string(output), nil
=======
	extractedText := string(output)
	log.Println("✅ Extracted Text:", extractedText)

	// Send extracted text to LLaMA for further processing using the interface
	analysis, err := llamaClient.AnalyzeMedication(context.Background(), extractedText)
	if err != nil {
		log.Printf("LLaMA Processing Failed: %v", err)
		return extractedText, nil // Return raw extracted text if LLaMA fails
	}

	log.Println("✅ Analyzed Medication:", analysis)
	// You can use the analysis data as needed or return it
	return fmt.Sprintf("Analysis Result: %v", analysis), nil
>>>>>>> 6b5b2b5ce4434bd8b80f655908dba12693114995
}
