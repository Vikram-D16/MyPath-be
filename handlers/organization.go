package handlers

import (
	"log"
	"net/http"

	"github.com/Vikram-D16/MyPath-be/db"
	"github.com/Vikram-D16/MyPath-be/models"
	"github.com/Vikram-D16/MyPath-be/ms/email"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateOrganizationHandler(c *gin.Context) {
	var org models.Organization

	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if org.Title == "" || org.Description == "" || org.Country == "" || org.State == "" || org.City == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	org.ID = uuid.New()
	if err := db.DB.Create(&org).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create organization"})
		return
	}

	// Send notification email
	if org.Email != nil && *org.Email != "" {
		if err := email.SendOrganizationEmail(*org.Email, org.Title); err != nil {
			log.Printf("Failed to send organization email: %v", err)
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Organization created successfully",
		"organization": org,
	})
}

func GetAllOrganizationsHandler(c *gin.Context) {
	var orgs []models.Organization

	query := db.DB.Model(&models.Organization{})

	if country := c.Query("country"); country != "" {
		query = query.Where("country = ?", country)
	}
	if state := c.Query("state"); state != "" {
		query = query.Where("state = ?", state)
	}
	if city := c.Query("city"); city != "" {
		query = query.Where("city = ?", city)
	}
	if isActive := c.Query("is_active"); isActive != "" {
		query = query.Where("is_active = ?", isActive)
	}
	if err := query.Where("is_deleted IS NULL").Find(&orgs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve organizations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"organizations": orgs})
}
