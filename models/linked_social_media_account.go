package models

type LinkedSocialMediaAccount struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	UserID         uint   `json:"user_id"`
	ProviderName   string `json:"provider_name"`
	ProviderID     string `json:"provider_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func (LinkedSocialMediaAccount) TableName() string {
	return "linked_social_accounts"
}
