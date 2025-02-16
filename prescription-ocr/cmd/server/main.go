package main

import (
	"github.com/Aanandvyas/Health_Hackathon/prescription-ocr/internal/api/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/upload", handlers.HandleUpload).Methods("POST")
	router.HandleFunc("/process", handlers.ProcessPrescription).Methods("POST")

	log.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
