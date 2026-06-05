package models

import (
	"time"

	"gorm.io/gorm"
)

// Model User merepresentasikan tabel "users" di database
type User struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	Password        string `json:"password"`
	StatusId        int    `json:"status_id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "users" // jadi singular
}
