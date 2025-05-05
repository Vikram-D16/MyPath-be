package handlers

import (
	"fmt"
	"net/http"

	"github.com/Vikram-D16/MyPath-be/db"
	"github.com/Vikram-D16/MyPath-be/models"
	"github.com/Vikram-D16/MyPath-be/utils"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields (username, email, password) are required"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	if err := models.RegisterUser(db.DB, user.Username, user.Email, hashedPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error registering user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the MYPATH Career Growth App!")
}

func GetAllUsersHandler(c *gin.Context) {
	users, err := models.GetAllUsers(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// LoginHandler handles user login
func LoginHandler(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if credentials.Email == "" || credentials.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Find user
	var user models.User
	if err := db.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Validate password
	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check for existing active token
	var userToken models.UserToken
	err := db.DB.
		Where("user_id = ? AND is_active = ?", user.ID, true).
		First(&userToken).Error

	if err != nil || userToken.Token == "" {
		newToken := utils.GenerateRandomToken()

		userToken = models.UserToken{
			UserID:   user.ID,
			Token:    newToken,
			IsActive: true,
		}

		if err := db.DB.Create(&userToken).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate user token"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"userName":    user.Username,
		"email":       user.Email,
		"message":     "Login successful",
		"accessToken": userToken.Token,
	})
}
