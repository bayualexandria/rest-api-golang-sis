package models

import (
	"time"

	"gorm.io/gorm"
)

type JadwalPelajaran struct {
	Id              int64     `json:"id"`
	KelasId         uint64    `json:"kelas_id"`
	GuruId          uint64    `json:"guru_id"`
	MataPelajaranId uint64    `json:"mata_pelajaran_id"`
	TahunAjaranId   uint64    `json:"tahun_ajaran_id"`
	Hari            string    `json:"hari"`
	JamMulai        time.Time `json:"jam_mulai"`
	JamSelesai      time.Time `json:"jam_selesai"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (JadwalPelajaran) TableName() string {
	return "jadwal_pelajaran"
}
