package ocr

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
	output, err := cmd.CombinedOutput()

	// Error handling
	if err != nil {
		log.Printf("Tesseract Error: %s\n", string(output))
		return "", fmt.Errorf("failed to process image: %s", string(output))
	}

	// Log the extracted text (optional)
	log.Println("✅ Extracted Text: ", string(output))

	// Return the extracted text
	return string(output), nil
}
