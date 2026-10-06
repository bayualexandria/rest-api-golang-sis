package models

import (
	"time"

	"gorm.io/gorm"
)

// Model WaliKelas merepresentasikan tabel "wali_kelas" di database
type WaliKelas struct {
	GuruWaliId    uint64 `json:"guru_wali_id"`
	KelasId       uint64 `json:"kelas_id"`
	TahunAjaranId uint64 `json:"tahun_ajaran_id"`
	SemesterId    uint64 `json:"semester_id"`
	Status        string `json:"status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (WaliKelas) TableName() string {
	return "wali_kelas" // jadi singular
}
