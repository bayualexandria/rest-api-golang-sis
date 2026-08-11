package models

import (
	"time"

	"gorm.io/gorm"
)

type SiswaKelas struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`

	SiswaID       uint64 `gorm:"not null;index" json:"siswa_id"`
	KelasID       uint64 `gorm:"not null;index" json:"kelas_id"`
	TahunAjaranID uint64 `gorm:"not null;index" json:"tahun_ajaran_id"`
	SemesterID     uint64 `gorm:"not null;index" json:"semester_id"`

	TanggalMasuk *time.Time `gorm:"type:date" json:"tanggal_masuk"`

	Status string `gorm:"size:50;not null;default:aktif" json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	// Siswa       Siswa       `gorm:"foreignKey:SiswaID;references:ID" json:"siswa"`
	// Kelas       Kelas       `gorm:"foreignKey:KelasID;references:ID" json:"kelas"`
	// TahunAjaran TahunAjaran `gorm:"foreignKey:TahunAjaranID;references:ID" json:"tahun_ajaran"`
}

func (SiswaKelas) TableName() string {
	return "siswa_kelas"
}