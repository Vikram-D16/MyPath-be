package models

import (
	"context"
	"fmt"

	"github.com/Vikram-D16/MyPath-be/db"
)

// RegisterUser inserts a new user into the database
func RegisterUser(username, email, passwordHash string) error {
	// Insert user into the 'users' table
	_, err := db.Conn.Exec(context.Background(),
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		username, email, passwordHash,
	)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}
	return nil
}
