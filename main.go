// main.go
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// User struct for capturing registration details from the request body
type User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Register handler to handle user registration
func registerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var user User
    // Decode the incoming JSON request body into the User struct
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate the user input
    if user.Username == "" || user.Email == "" || user.Password == "" {
        http.Error(w, "All fields (username, email, password) are required", http.StatusBadRequest)
        return
    }

    // Hash the password before storing it
    hashedPassword := hashPassword(user.Password)

    // Call RegisterUser function to insert the user into the database
    err = RegisterUser(user.Username, user.Email, hashedPassword)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error registering user: %v", err), http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintln(w, "User registered successfully!")
}

// Hash the password using SHA256 (Consider using bcrypt for production)
func hashPassword(password string) string {
    hash := sha256.New()
    hash.Write([]byte(password))
    return hex.EncodeToString(hash.Sum(nil))
}

// Home handler to return a simple welcome message
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Career Growth App!")
}

func main() {
    // Initialize the database connection
    InitDB()
    defer CloseDB()

    // Set up routes
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/register", registerHandler)

    // Start the server on port 8080
    fmt.Println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal("Error starting the server:", err)
    }
}
