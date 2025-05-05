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
