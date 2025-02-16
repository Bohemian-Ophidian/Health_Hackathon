package ocr

import (
	"log"
	"os/exec"
)

func ExtractTextFromImage(imagePath string) string {
	cmd := exec.Command("tesseract", imagePath, "stdout")

	// Capture the output
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Return the extracted text
	return string(output)
}
