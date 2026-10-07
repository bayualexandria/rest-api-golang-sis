package models

import "time"

type GuruMapel struct {
	Id              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	GuruId          uint64    `gorm:"not null;index" json:"guru_id"`
	MataPelajaranId uint64    `gorm:"not null;index" json:"mata_pelajaran_id"`
	TahunAjaranId   uint64    `gorm:"not null;index" json:"tahun_ajaran_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (GuruMapel) TableName() string {
	return "guru_mata_pelajaran"
}
