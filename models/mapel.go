package models

import (
	"time"

	"gorm.io/gorm"
)

type Mapel struct {
	Id        uint64 `json:"id"`
	Nama      string `json:"nama"`
	Kode      string `json:"kode"`
	Deskripsi string `json:"deskripsi"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Mapel) TableName() string {
	return "mata_pelajaran"
}
