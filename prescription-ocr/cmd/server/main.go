package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/otiai10/gosseract"
)

var (
	huggingFaceAPIUrl = ""
	huggingFaceAPIKey = ""
	tesseractPath     = ""
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load the environment variables at runtime (not in constants)
	huggingFaceAPIUrl = os.Getenv("HUGGINGFACE_API_URL")
	huggingFaceAPIKey = os.Getenv("HUGGINGFACE_API_KEY")
	tesseractPath = os.Getenv("TESSERACT_PATH")
}

func extractTextFromImage(imagePath string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	// Set the Tesseract executable path if required in your system
	client.SetImage(imagePath)

	// Perform OCR
	extractedText, err := client.Text()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(extractedText), nil
}

func queryModel(medicineName string) (string, error) {
	// Construct the prompt for the model
	prompt := fmt.Sprintf(`
	Please provide information about the following medicine in two sections:
	1. **Alternative Medicines**: List other medicines that can be used for the same purpose.
	2. **Side Effects**: List common side effects of %s.

	Medicine Name: %s
	`, medicineName, medicineName)

	// Create the request body
	requestBody := fmt.Sprintf(`{"inputs": "%s"}`, prompt)

	// Make the HTTP request to the Hugging Face model API
	req, err := http.NewRequest("POST", huggingFaceAPIUrl, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		return "", err
	}

	// Set necessary headers
	req.Header.Set("Authorization", "Bearer "+huggingFaceAPIKey)
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert response to string and return
	return string(body), nil
}

func extractMedicineNameHandler(w http.ResponseWriter, r *http.Request) {
	// Expecting an image file uploaded via multipart form
	r.ParseMultipartForm(10 << 20) // Limit file size to 10MB

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file temporarily
	tempFilePath := "./temp_image.jpg"
	outFile, err := ioutil.TempFile("", "temp_image_*.jpg")
	if err != nil {
		http.Error(w, "Error saving temporary image", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading image data", http.StatusInternalServerError)
		return
	}

	// Extract text from the image
	medicineName, err := extractTextFromImage(tempFilePath)
	if err != nil || medicineName == "" {
		http.Error(w, "Failed to extract text or medicine name is empty", http.StatusInternalServerError)
		return
	}

	// Return the extracted medicine name as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"medicine_name": medicineName,
	})
}

func queryMedicineInfoHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		MedicineName string `json:"medicine_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Get model response
	modelResponse, err := queryModel(requestData.MedicineName)
	if err != nil {
		http.Error(w, "Error querying the model", http.StatusInternalServerError)
		return
	}

	// Return model response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"model_response": modelResponse,
	})
}

func main() {
	http.HandleFunc("/extract_medicine_name", extractMedicineNameHandler)
	http.HandleFunc("/query_medicine_info", queryMedicineInfoHandler)

	// Start the HTTP server on port 8080
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
