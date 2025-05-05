package models

type UserToken struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserID   uint   `gorm:"not null;unique" json:"user_id"`
	Token    string `gorm:"not null;unique" json:"token"`
	IsActive bool   `gorm:"not null" json:"is_active"`
}
