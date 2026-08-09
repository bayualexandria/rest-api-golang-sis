package seeders

import (
	"log"

	"gorm.io/gorm"
)

type StatusSiswaSeeder struct {
	NamaStatus string
	Keterangan string
}

func (StatusSiswaSeeder) TableName() string {
	return "status_siswa" // jadi singular
}

func (s StatusSiswaSeeder) Run(db *gorm.DB) {
	statusSiswaList := []StatusSiswaSeeder{
		{NamaStatus: "Aktif", Keterangan: "Siswa sedang aktif"},
		{NamaStatus: "Tidak Aktif", Keterangan: "Siswa tidak aktif"},
		{NamaStatus: "Lulus", Keterangan: "Siswa telah lulus"},
		{NamaStatus: "Dikeluarkan", Keterangan: "Siswa keluar sebelum lulus"},
		{NamaStatus: "Pindah", Keterangan: "Siswa pindah ke sekolah lain"},
	}
	for _, status := range statusSiswaList {
		if err := db.Create(&status).Error; err != nil {
			log.Fatal("Error creating status siswa:", err)
		}
	}
	log.Println("Seeder Status Siswa selesai")
}
