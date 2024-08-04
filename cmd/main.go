package main

import (
	"chat-app/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Create a new HTTP server
	router := handlers.SetupRouter()

	// Start the server
	log.Println("Starting chat server on : 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server : %v", err)
	}
}
