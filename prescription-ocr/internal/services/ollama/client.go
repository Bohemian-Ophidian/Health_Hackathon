package ollama

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func ProcessPrescription(text string) string {
	// Updated to use the correct Ollama API endpoint (localhost:11434)
	url := "http://localhost:11434/process" // Use the correct port

	// Prepare the JSON request body
	requestBody, err := json.Marshal(map[string]string{
		"text": text,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Make the POST request to Ollama API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse the response
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	// Assuming the response contains medication information
	return result["medicationDetails"].(string)
}
