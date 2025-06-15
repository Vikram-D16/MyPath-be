package models

import (
	"time"

	"github.com/google/uuid"
)

type GroupMember struct {
	GroupID   uuid.UUID  `gorm:"type:uuid;primaryKey" json:"group_id"`
	UserID    uuid.UUID  `gorm:"type:uuid;primaryKey" json:"user_id"`
	RoleID    int        `gorm:"not null" json:"role_id"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
