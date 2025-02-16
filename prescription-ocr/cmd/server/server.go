package main

import (
	"log"
	"net/http"
)

func setupRoutes() {
	// Route mappings can go here if needed.
	http.HandleFunc("/", handleHome)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Prescription OCR API"))
}

func StartServer() {
	setupRoutes()
	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
