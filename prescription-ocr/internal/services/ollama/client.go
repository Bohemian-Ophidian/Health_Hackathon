package ollama

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func ProcessPrescription(text string) string {
	// Ollama API endpoint
	url := "http://ollama_api_endpoint" // Replace with actual Ollama API URL

	requestBody, err := json.Marshal(map[string]string{
		"text": text,
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	// Assuming the response contains the medication details
	return result["medicationDetails"].(string) // Adjust based on the actual API response
}
