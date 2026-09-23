package models

import (
	"time"

	"gorm.io/gorm"
)

type Kelas struct {
	Id        int64  `json:"id"`
	NamaKelas string `json:"nama_kelas"`
	Jurusan   string `json:"jurusan"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Kelas) TableName() string {
	return "kelas"
}
