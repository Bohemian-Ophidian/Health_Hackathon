package ocr

import (
	"log"

	"github.com/otiai10/gosseract"
)

func ExtractTextFromImage(imagePath string) string {
	client := gosseract.NewClient()
	defer client.Close()

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
