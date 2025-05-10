package handlers

import (
	"fmt"
	"net/http"

	"github.com/Vikram-D16/MyPath-be/db"
	"github.com/Vikram-D16/MyPath-be/models"
	"github.com/Vikram-D16/MyPath-be/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var missingFields []string
	if user.Username == "" {
		missingFields = append(missingFields, "username")
	}
	if user.Email == "" {
		missingFields = append(missingFields, "email")
	}
	if user.Password == "" {
		missingFields = append(missingFields, "password")
	}

	if len(missingFields) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Missing fields: %v", missingFields)})
		return
	}

	var existingUser models.User
	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	hashedPwd, err := utils.GeneratePassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	if err := utils.RegisterUser(db.DB, user.Username, user.Email, hashedPwd, user.Avatar, user.Bio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error registering user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": gin.H{
		"id": user.ID, "username": user.Username, "email": user.Email, "avatar": user.Avatar, "bio": user.Bio}})
}

func LoginHandler(c *gin.Context) {
	var creds struct {
		Email, Password string
	}
	if err := c.ShouldBindJSON(&creds); err != nil || creds.Email == "" || creds.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", creds.Email).First(&user).Error; err != nil ||
		!utils.CheckPasswordHash(creds.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	var token models.UserToken
	if err := db.DB.Where("user_id = ? AND is_active = true", user.ID).First(&token).Error; err != nil {
		token = models.UserToken{UserID: user.ID, Token: utils.GenerateRandomToken(), IsActive: true}
		db.DB.Create(&token)
	}

	c.JSON(http.StatusOK, gin.H{
		"userName": user.Username, "email": user.Email,
		"message": "Login successful", "accessToken": token.Token,
	})
}

func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the MYPATH Career Growth App!")
}

func GetAllUsersHandler(c *gin.Context) {
	users, err := utils.GetAllUsers(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func DeleteUserHandler(c *gin.Context) {
	userIdInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIdInterface.(uuid.UUID)
	fmt.Println(userID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	// Delete user record
	if err := db.DB.Delete(&models.User{}, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Delete any active tokens for this user
	db.DB.Where("user_id = ?", userID).Delete(&models.UserToken{})

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
