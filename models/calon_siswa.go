package models

import "time"

type CalonSiswa struct {
	Id               uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	IdUserCalonSiswa string    `gorm:"not null;index" json:"id_user_calon_siswa"`
	Nisn             string    `gorm:"not null;index" json:"nisn"`
	Nik              string    `gorm:"not null;index" json:"nik"`
	NamaLengkap      string    `gorm:"not null" json:"nama_lengkap"`
	Ttl              string    `gorm:"not null" json:"ttl"`
	JenisKelamin     string    `gorm:"not null" json:"jenis_kelamin"`
	Agama            string    `gorm:"not null" json:"agama"`
	Alamat           string    `gorm:"not null" json:"alamat"`
	SekolahAsal      string    `gorm:"not null" json:"sekolah_asal"`
	TahunLulus       string    `gorm:"not null" json:"tahun_lulus"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (CalonSiswa) TableName() string {
	return "calon_siswa"
}
