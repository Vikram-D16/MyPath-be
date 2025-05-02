package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Vikram-D16/MyPath-be/models"
	"github.com/Vikram-D16/MyPath-be/utils"
)

// RegisterHandler handles the user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Decode the incoming JSON request body into the User struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate user input
	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "All fields (username, email, password) are required", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Insert user into the database
	if err := models.RegisterUser(user.Username, user.Email, hashedPassword); err != nil {
		http.Error(w, fmt.Sprintf("Error registering user: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User registered successfully!")
}

// HomeHandler returns a welcome message
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the MYPATH Career Growth App!")
}
