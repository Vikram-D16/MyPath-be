package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Group struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null" json:"organization_id"`
	Title          string         `gorm:"not null" json:"title"`
	Description    string         `gorm:"type:text;not null" json:"description"`
	Info           datatypes.JSON `json:"info" gorm:"type:jsonb;default:'{}'"`
	IsActive       string         `gorm:"type:enum('active','inactive');default:'inactive'" json:"is_active"`
	IsDeleted      *time.Time     `json:"is_deleted,omitempty"`
	Picture        *int           `json:"picture,omitempty"`
	MembersCount   int            `gorm:"default:0" json:"members_count"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
