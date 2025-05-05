package utils

import (
	"github.com/Vikram-D16/MyPath-be/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RegisterUser inserts a new user into the database
func RegisterUser(db *gorm.DB, username, email, passwordHash string) error {
	user := models.User{
		ID:       uuid.New(),
		Username: username,
		Email:    email,
		Password: passwordHash,
	}
	return db.Create(&user).Error
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}
