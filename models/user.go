package models

// Model User merepresentasikan tabel "users" di database
type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Username        int    `json:"username"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	Password        string `json:"password"`
	StatusUserId        string `json:"status_user_id"`
}

func (User) TableName() string {
	return "users" // jadi singular
}
