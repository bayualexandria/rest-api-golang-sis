package models

import "time"

type LinkedSocialMediaAccount struct {
	Id           uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       string `json:"user_id"`
	ProviderName string `json:"provider_name"`
	ProviderID   string `json:"provider_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (LinkedSocialMediaAccount) TableName() string {
	return "linked_social_accounts"
}
