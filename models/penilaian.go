package models

import "time"

type Penilaian struct {
	Id               uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	KelasId          uint64    `gorm:"not null;index" json:"kelas_id"`
	MataPelajaranId  uint64    `gorm:"not null;index" json:"mata_pelajaran_id"`
	GuruId           uint64    `gorm:"not null;index" json:"guru_id"`
	JenisPenilaianId uint64    `gorm:"not null;index" json:"jenis_penilaian_id"`
	SemesterId       uint64    `gorm:"not null;index" json:"semester_id"`
	NamaPenilaian    string    `gorm:"not null;index" json:"nama_penilaian"`
	Tanggal          time.Time `gorm:"not null;index" json:"tanggal"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (Penilaian) TableName() string {
	return "penilaian"
}
