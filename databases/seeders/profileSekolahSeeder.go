package seeders

import (
	"log"

	"gorm.io/gorm"
)

type ProfileSekolahSeeder struct {
	ID           int
	NamaSekolah  string
	Alamat       string
	NoTelp       string
	Akreditasi   string
	ImageProfile string
}

func (ProfileSekolahSeeder) TableName() string {
	return "profile_sekolah" // jadi singular
}

func (p ProfileSekolahSeeder) Run(db *gorm.DB) {
	profile := ProfileSekolahSeeder{
		ID:           1,
		NamaSekolah:  "SMK NEGERI 1 SINGOSARI",
		Alamat:       "Jl. Raya Singosari No. 1, Singosari, Malang",
		NoTelp:       "0341-123456",
		Akreditasi:   "A",
		ImageProfile: "storages/logo-pendidikan.png",
	}
	if err := db.Create(&profile).Error; err != nil {
		log.Fatal("Error creating profile:", err)
	}

	log.Println("Seeder Profile Sekolah selesai")
}
