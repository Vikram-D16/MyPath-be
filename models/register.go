package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RegisterUser inserts a new user into the database
func RegisterUser(db *gorm.DB, username, email, passwordHash string) error {
	user := User{
		ID:       uuid.New(),
		Username: username,
		Email:    email,
		Password: passwordHash,
	}
	return db.Create(&user).Error
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}
