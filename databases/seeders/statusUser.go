package seeders

import (
	"log"

	"gorm.io/gorm"
)

type StatusUserSeeder struct {
	ID     int
	NamaStatus string
}

func (StatusUserSeeder) TableName() string {
	return "status_user" // jadi singular
}

func (s StatusUserSeeder) Run(db *gorm.DB) {
	statusUser := []StatusUserSeeder{
		{ID: 1, NamaStatus: "Admin"},
		{ID: 2, NamaStatus: "Wali Kelas"},
		{ID: 3, NamaStatus: "Guru"},
		{ID: 4, NamaStatus: "Siswa"},
	}
	for _, su := range statusUser {
		if err := db.Create(&su).Error; err != nil {
			log.Println("Gagal insert:", err)
		}
	}

	log.Println("Seeder StatusUser selesai")
}
