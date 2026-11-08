package models

import "time"

type OrangTua struct {
	Id            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	IdSiswa       string    `gorm:"not null;index" json:"id_siswa"`
	NikAyah       string    `gorm:"not null;index" json:"nik_ayah"`
	NamaAyah      string    `gorm:"not null" json:"nama_ayah"`
	PekerjaanAyah string    `gorm:"not null" json:"pekerjaan_ayah"`
	NikIbu        string    `gorm:"not null;index" json:"nik_ibu"`
	NamaIbu       string    `gorm:"not null" json:"nama_ibu"`
	PekerjaanIbu  string    `gorm:"not null" json:"pekerjaan_ibu"`
	NamaWali      string    `gorm:"not null" json:"nama_wali"`
	PekerjaanWali string    `gorm:"not null" json:"pekerjaan_wali"`
	NomorTelepon  string    `gorm:"not null" json:"nomor_telepon"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (OrangTua) TableName() string {
	return "orang_tua"
}
