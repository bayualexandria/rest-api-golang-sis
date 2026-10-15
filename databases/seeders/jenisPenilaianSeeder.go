package seeders

import (
	"log"

	"gorm.io/gorm"
)

type JenisPenilaianSeeder struct {
	Nama string
	Kode string
}

func (JenisPenilaianSeeder) TableName() string {
	return "jenis_penilaian" // jadi singular
}

func (j JenisPenilaianSeeder) Run(db *gorm.DB) {
	jenisPenilaianList := []JenisPenilaianSeeder{
		{Nama: "Penilaian Harian", Kode: "PH"},
		{Nama: "Penilaian Tengah Semester", Kode: "PTS"},
		{Nama: "Penilaian Akhir Semester", Kode: "PAS"},
		{Nama: "Penilaian Tugas Rumah", Kode: "PTR"},
		{Nama: "Penilaian Praktik / Unjuk Kerja", Kode: "PRK"},
		{Nama: "Ujian Sekolah", Kode: "US"},
	}
	for _, jenisPenilaian := range jenisPenilaianList {
		if err := db.Create(&jenisPenilaian).Error; err != nil {
			log.Fatal("Error creating jenis penilaian:", err)
		}
	}

	log.Println("Seeder Jenis Penilaian selesai")
}
