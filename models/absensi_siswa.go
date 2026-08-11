package models

import (
	"time"

)

type AbsensiSiswa struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`

	SiswaKelasID uint64 `gorm:"not null;index" json:"siswa_kelas_id"`

	SemesterID uint64 `gorm:"not null;index" json:"semester_id"`

	Tanggal time.Time `gorm:"type:date;not null;index" json:"tanggal"`

	StatusKehadiranID uint64 `gorm:"not null;index" json:"status_kehadiran_id"`

	Keterangan *string `gorm:"type:text" json:"keterangan"`

	JamMasuk *time.Time `gorm:"type:time" json:"jam_masuk"`

	JamKeluar *time.Time `gorm:"type:time" json:"jam_keluar"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`

}

func (AbsensiSiswa) TableName() string {
	return "absensi_siswa"
}
