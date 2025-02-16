package utils

import "log"

// Example utility function for logging
func LogError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
