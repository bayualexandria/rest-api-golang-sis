package models

// Model PersonalAccessToken merepresentasikan tabel "personal_access_tokens" di database
type PersonalAccessToken struct {
	Token         string `json:"token"`
	TokenableType string `json:"tokenable_type"`
	TokenableID   string    `json:"tokenable_id"`
	Name          string `json:"name"`
	Abilities     string `json:"abilities"`
	LastUsedAt    string `json:"last_used_at"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (PersonalAccessToken) TableName() string {
	return "personal_access_tokens" // jadi singular
}
