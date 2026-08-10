package models

import (
	"time"

	"gorm.io/gorm"
)

type TahunAjaran struct {
	ID             uint64         `gorm:"primaryKey;column:id" json:"id"`
	NamaTahun      string         `gorm:"column:nama_tahun;size:20;not null" json:"nama_tahun"`
	TanggalMulai   time.Time      `gorm:"column:tanggal_mulai;not null" json:"tanggal_mulai"`
	TanggalSelesai time.Time      `gorm:"column:tanggal_selesai;not null" json:"tanggal_selesai"`
	IsActive       bool           `gorm:"column:is_active;not null;default:false" json:"is_active"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (TahunAjaran) TableName() string {
	return "tahun_ajaran" // jadi singular
}
