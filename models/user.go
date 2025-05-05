package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	Username  string          `gorm:"unique;not null" json:"username"`
	Email     string          `gorm:"unique;not null" json:"email"`
	Password  string          `gorm:"not null" json:"password"`
	CreatedAt *gorm.DeletedAt `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *gorm.DeletedAt `gorm:"autoUpdateTime" json:"updated_at"`
}
