package seeders

import (
	"log"

	"gorm.io/gorm"
)

type StatusKehadiranSeeder struct {
	NamaStatus string
	Kode 	 string
}

func (StatusKehadiranSeeder) TableName() string {
	return "status_kehadiran" // jadi singular
}

func (s StatusKehadiranSeeder) Run(db *gorm.DB) {
	statusKehadiranList := []StatusKehadiranSeeder{
		{NamaStatus: "Hadir", Kode: "H"},
		{NamaStatus: "Sakit", Kode: "S"},
		{NamaStatus: "Izin", Kode: "I"},
		{NamaStatus: "Alpa", Kode: "A"},
	}
	for _, status := range statusKehadiranList {
		if err := db.Create(&status).Error; err != nil {
			log.Fatal("Error creating status kehadiran:", err)
		}
	}
	log.Println("Seeder Status Kehadiran selesai")
}
