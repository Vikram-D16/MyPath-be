package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserToken struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	UserID    uuid.UUID       `gorm:"not null;unique" json:"user_id"`
	Token     string          `gorm:"not null;unique" json:"token"`
	IsActive  bool            `gorm:"not null" json:"is_active"`
	CreatedAt *gorm.DeletedAt `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *gorm.DeletedAt `gorm:"autoUpdateTime" json:"updated_at"`
}
