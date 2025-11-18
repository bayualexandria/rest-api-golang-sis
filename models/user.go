package models

// Model User merepresentasikan tabel "users" di database
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username int    `json:"username"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "users" // jadi singular
}
