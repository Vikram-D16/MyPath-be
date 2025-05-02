package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/Vikram-D16/MyPath-be/db"
	"github.com/Vikram-D16/MyPath-be/handlers"
)

func main() {
	// Initialize the database connection
	db.Init()
	defer db.Close()

	// Set up the HTTP router
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

	// Get the server port from the environment variables or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Start the HTTP server
	fmt.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
