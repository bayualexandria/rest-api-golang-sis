package models

import "time"

type Pendaftaran struct {
	Id                 uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	NomorPendaftaran   string    `gorm:"not null;index" json:"nomor_pendaftaran"`
	IdJalurPendaftaran int       `gorm:"not null;index" json:"id_jalur_pendaftaran"`
	TanggalDaftar      string    `gorm:"not null;index" json:"tanggal_daftar"`
	StatusVerifikasi   string    `gorm:"not null" json:"status_verifikasi"`
	CatatanVerifikasi  string    `gorm:"not null" json:"catatan_verifikasi"`
	StatusKelulusan    string    `gorm:"not null" json:"status_kelulusan"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (Pendaftaran) TableName() string {
	return "pendaftaran"
}
