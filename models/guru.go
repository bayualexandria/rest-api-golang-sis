package models

import (
	"time"

	"gorm.io/gorm"
)

type Guru struct {
	Nip          int    `json:"nip"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHp         string `json:"no_hp"`
	Alamat       string `json:"alamat"`
	ImageProfile string `json:"image_profile"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (Guru) TableName() string {
	return "guru" // jadi singular
}
