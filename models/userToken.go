package models

import "gorm.io/gorm"

type UserToken struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	UserID    uint            `gorm:"not null;unique" json:"user_id"`
	Token     string          `gorm:"not null;unique" json:"token"`
	IsActive  bool            `gorm:"not null" json:"is_active"`
	CreatedAt *gorm.DeletedAt `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *gorm.DeletedAt `gorm:"autoUpdateTime" json:"updated_at"`
}
