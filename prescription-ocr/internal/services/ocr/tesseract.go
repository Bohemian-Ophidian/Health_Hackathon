package ocr

import (
	"log"

	"github.com/otiai10/gosseract"
)

func ExtractTextFromImage(imagePath string) string {
	client := gosseract.NewClient()
	defer client.Close()

	// Set the path to the Tesseract executable if it's not in your system PATH
	client.SetTesseractPath("C:\\Program Files\\Tesseract-OCR\\tesseract.exe") // Adjust the path if necessary

	err := client.SetImage(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
	}
	return text
}
