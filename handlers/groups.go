package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Vikram-D16/MyPath-be/db"
	"github.com/Vikram-D16/MyPath-be/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateGroupHandler(c *gin.Context) {
	var group models.Group

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if group.Title == "" || group.Description == "" || group.OrganizationID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	var org models.Organization
	if err := db.DB.First(&org, "id = ?", group.OrganizationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	group.ID = uuid.New()
	if err := db.DB.Create(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Group created successfully", "group": group})
}

func CreateGroupMemberHandler(c *gin.Context) {
	var groupMember models.GroupMember

	if err := c.ShouldBindJSON(&groupMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	if groupMember.GroupID == uuid.Nil || groupMember.UserID == uuid.Nil || groupMember.RoleID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	var grp models.Group
	if err := db.DB.First(&grp, "id = ?", groupMember.GroupID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}

	var usr models.User
	if err := db.DB.First(&usr, "id = ?", groupMember.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.DB.Create(&groupMember).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group member"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Group Member created successfully", "group": groupMember})

}

func GetAllGroupsHandler(c *gin.Context) {
	var groups []models.Group

	if err := db.DB.Where("is_deleted IS NULL AND is_active = ?", "active").Find(&groups).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve groups"})
		return
	}
	fmt.Print(groups)
	c.JSON(http.StatusOK, groups)
}

func DeleteGroupHandler(c *gin.Context) {
	groupIDParam := c.Param("id")
	groupID, err := uuid.Parse(groupIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}
	fmt.Print(groupID)
	var group models.Group
	if err := db.DB.Where("id = ? AND is_deleted IS NULL AND is_active = ?", groupID, "active").First(&group).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found or already deleted"})
		return
	}

	now := time.Now()
	group.IsDeleted = &now

	if err := db.DB.Save(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete group"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
}
