// db.go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    "github.com/jackc/pgx/v4"
)

var db *pgx.Conn

// Initialize the DB connection
func InitDB() {
    var err error
    dsn := "postgres://postgres:yourpassword@localhost:5432/mydb"

    // Retry logic to connect to the database
    for i := 0; i < 10; i++ {  // Retry 10 times with a delay 
        db, err = pgx.Connect(context.Background(), dsn)
        if err == nil {
            fmt.Println("✅ Connected to PostgreSQL DB")
            return
        }

        log.Printf("Unable to connect to database, retrying in 2 seconds... (Attempt %d/10)\n", i+1)
        time.Sleep(2 * time.Second)  // Sleep for 2 seconds before retrying
    }

    if err != nil {
        log.Fatalf("Unable to connect to database after 10 retries: %v\n", err)
    }
}

// CloseDB function to close the database connection
func CloseDB() {
    if db != nil {
        db.Close(context.Background())
    }
}


// Function to insert a new user into the database
func RegisterUser(username, email, passwordHash string) error {
	fmt.Println("Database connection inserted.")
	fmt.Println("Connection is newly started.")
	
    // Insert user into the 'users' table
    _, err := db.Exec(context.Background(), "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, passwordHash)
    if err != nil {
        return fmt.Errorf("Error inserting user into database: %v", err)
    }
    return nil
}
