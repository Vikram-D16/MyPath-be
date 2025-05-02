package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
"github.com/joho/godotenv"
	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

// Init initializes the database connection with retry logic
func Init() {
	var err error
	// Fetch database URL from environment variable
	t := godotenv.Load()

	if t != nil {
		fmt.Println("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	// Retry logic to connect to the database
	for i := 0; i < 10; i++ {
		Conn, err = pgx.Connect(context.Background(), dsn)
		if err == nil {
			fmt.Println("✅ Connected to PostgreSQL DB")
			return
		}
		log.Printf("Unable to connect to database, retrying in 2 seconds... (Attempt %d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	// If still unable to connect after 10 attempts, log fatal error
	log.Fatalf("Unable to connect to database after 10 retries: %v\n", err)
}

// Close closes the database connection
func Close() {
	if Conn != nil {
		Conn.Close(context.Background())
	}
}
