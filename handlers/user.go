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
	if err := c.ShouldBindJSON(&user); err != nil ||
		user.Username == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields (username, email, password) are required"})
		return
	}

	hashedPwd, err := utils.GeneratePassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	if err := models.RegisterUser(db.DB, user.Username, user.Email, hashedPwd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error registering user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
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
	// userId := c.MustGet("userId").(uint)
	users, err := models.GetAllUsers(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
