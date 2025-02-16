package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		URI    string
		DBName string
	}
	Server struct {
		Port string
	}
	OCR struct {
		TesseractPath string
	}
	LLaMA struct {
		APIURL string
		APIKey string
	}
}

func LoadConfig() (*Config, error) {
	// Load .env file (if present)
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env file not found: %v\n", err)
	}

	config := &Config{}

	// MongoDB URI and Database Name configuration
	config.Database.URI = getEnv("MONGO_URI", "mongodb://localhost:27017")
	config.Database.DBName = getEnv("DB_NAME", "healthDB")

	// Server Port configuration
	config.Server.Port = getEnv("SERVER_PORT", "8080")

	// OCR configuration
	config.OCR.TesseractPath = getEnv("TESSERACT_PATH", getDefaultTesseractPath())

	// LLaMA API configuration
	config.LLaMA.APIURL = getEnv("LLAMA_API_URL", "")
	config.LLaMA.APIKey = getEnv("LLAMA_API_KEY", "")

	// Check if LLaMA API URL is provided
	if config.LLaMA.APIURL == "" {
		return nil, fmt.Errorf("LLAMA_API_URL is required")
	}

	// Additional error checks for MongoDB URI
	if config.Database.URI == "" {
		return nil, fmt.Errorf("MONGO_URI is required")
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getDefaultTesseractPath returns the default Tesseract executable path based on the system's OS
func getDefaultTesseractPath() string {
	if runtime.GOOS == "windows" {
		// Set default path for Windows
		return "C:\\Program Files\\Tesseract-OCR\\tesseract.exe"
	}
	// Default for UNIX-based systems
	return "/usr/bin/tesseract"
}
