package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationMember struct {
	OrganizationID uuid.UUID      `gorm:"type:uuid;primaryKey" json:"organization_id"`
	UserID         uuid.UUID      `gorm:"type:uuid;primaryKey" json:"user_id"`
	RoleID         int            `gorm:"not null" json:"role_id"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
