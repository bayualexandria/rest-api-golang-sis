package seeders

import (
	"log"

	"gorm.io/gorm"
)

type StatusUserSeeder struct {
	NamaStatus string
}

func (StatusUserSeeder) TableName() string {
	return "status_user" // jadi singular
}

func (s StatusUserSeeder) Run(db *gorm.DB) {
	statusUser := []StatusUserSeeder{
		{NamaStatus: "Admin"},
		{NamaStatus: "Wali Kelas"},
		{NamaStatus: "Guru"},
		{NamaStatus: "Siswa"},
	}
	for _, su := range statusUser {
		if err := db.Create(&su).Error; err != nil {
			log.Println("Gagal insert:", err)
		}
	}

	log.Println("Seeder StatusUser selesai")
}
