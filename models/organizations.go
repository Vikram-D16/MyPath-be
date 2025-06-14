package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Organization struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CategoryID     *int           `json:"category_id,omitempty"`
	Title          string         `gorm:"not null" json:"title"`
	Description    string         `gorm:"type:text;not null" json:"description"`
	Address        *string        `json:"address,omitempty"`
	Country        string         `gorm:"not null" json:"country"`
	State          string         `gorm:"not null" json:"state"`
	City           string         `gorm:"not null" json:"city"`
	Zipcode        *string        `gorm:"size:255" json:"zipcode,omitempty"`
	Email          *string        `json:"email,omitempty"`
	Website        *string        `json:"website,omitempty"`
	Phone          *string        `json:"phone,omitempty"`
	About          *string        `gorm:"type:text" json:"about,omitempty"`
	IsActive       string         `gorm:"type:enum('active','inactive');default:'inactive'" json:"is_active"`
	IsDeleted      *time.Time     `json:"is_deleted,omitempty"`
	Picture        *int           `json:"picture,omitempty"`
	FollowersCount int            `gorm:"default:0" json:"followers_count"`
	MembersCount   int            `gorm:"default:0" json:"members_count"`
	Settings       datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"settings"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
