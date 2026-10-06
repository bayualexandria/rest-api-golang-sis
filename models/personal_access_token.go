package models

import "time"

// Model PersonalAccessToken merepresentasikan tabel "personal_access_tokens" di database
type PersonalAccessToken struct {
	Token         string    `json:"token"`
	TokenableType string    `json:"tokenable_type"`
	TokenableID   string    `json:"tokenable_id"`
	Name          string    `json:"name"`
	Abilities     string    `json:"abilities"`
	LastUsedAt    time.Time `json:"last_used_at"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (PersonalAccessToken) TableName() string {
	return "personal_access_tokens" // jadi singular
}
