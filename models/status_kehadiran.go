package models

import (
	"time"

)

type StatusKehadiran struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`

	NamaStatus string `gorm:"not null;index" json:"nama_status"`

	Kode string `gorm:"not null;index" json:"semester_id"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`

}

func (StatusKehadiran) TableName() string {
	return "status_kehadiran"
}
