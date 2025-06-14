package models

import (
	"time"

	"github.com/google/uuid"
)

type UserToken struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"not null;unique" json:"user_id"`
	Token     string    `gorm:"not null;unique" json:"token"`
	IsActive  bool      `gorm:"not null" json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
