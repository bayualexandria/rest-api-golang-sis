package models
type PasswordResetToken struct {
	Email     string `gorm:"type:varchar(100);not null" json:"email"`
	Token     string `gorm:"type:varchar(255);not null" json:"token"`
	CreatedAt string `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
}
func (PasswordResetToken) TableName() string {
	return "password_reset_tokens"
}