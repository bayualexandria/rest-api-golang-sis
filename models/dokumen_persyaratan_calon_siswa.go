package models

import "time"

type DokumenPersyaratanCalonSiswa struct {
	Id               uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	IdUserCalonSiswa string    `gorm:"not null;index" json:"id_user_calon_siswa"`
	JenisDokumen     string    `gorm:"not null" json:"jenis_dokumen"`
	NamaFile         string    `gorm:"not null" json:"nama_file"`
	PathFile         string    `gorm:"not null" json:"path_file"`
	DiunggahPada     time.Time `gorm:"not null" json:"diunggah_pada"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (DokumenPersyaratanCalonSiswa) TableName() string {
	return "dokumen_persyaratan_calon_siswa"
}
