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

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]User, error) {
	rows, err := db.Conn.Query(context.Background(), "SELECT username, email FROM users")
	if err != nil {
		return nil, fmt.Errorf("error querying users: %w", err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Username, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
