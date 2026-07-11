package seeders

import (
	"backend-api/utils"
	"log"
	"time"

	"gorm.io/gorm"
)

type GuruSeeder struct {
	Nip          string
	Nama         string
	JenisKelamin string
	NoHp         string
	Alamat       string
	ImageProfile string
}

type UserGuruSeeder struct {
	Name            string
	Username        string
	Email           string
	Password        string
	EmailVerifiedAt string
	StatusID        string
}

func (GuruSeeder) TableName() string {
	return "guru" // jadi singular
}

func (UserGuruSeeder) TableName() string {
	return "users" // jadi singular
}

func (s GuruSeeder) Run(db *gorm.DB) {

	guru := []GuruSeeder{
		{

			Nip:          "9106012508950001", // Contoh NIP dengan 8 digit
			Nama:         "Bayu Wardana",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567890",
			Alamat:       "Jl. Contoh Alamat No. 123, Kota Contoh",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{
			Nip:          "9106012508950002", // Contoh NIP dengan 8 digit
			Nama:         "Bayu Wardana 1",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567891",
			Alamat:       "Jl. Contoh Alamat No. 123, Kota Contoh",
			ImageProfile: "storage/logo-pendidikan.png",
		},
		{

			Nip:          "9106012508950003", // Contoh NIP dengan 8 digit
			Nama:         "Bayu Wardana 2",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567892",
			Alamat:       "Jl. Contoh Alamat No. 123, Kota Contoh",
			ImageProfile: "storage/logo-pendidikan.png",
		},
	}
	passHash := "admin123" // Contoh password default
	user := []UserGuruSeeder{
		{
			Name:            "Bayu Wardana",
			Username:        "9106012508950001",
			Email:           "wardanabayu453@gmail.com",
			Password:        utils.HashPasswordUser(passHash),
			EmailVerifiedAt: time.Now().Format("2006-01-02"),
			StatusID:        "1", // Misalnya, ID status user untuk Guru
		},
		{
			Name:            "Bayu Wardana 1",
			Username:        "9106012508950002",
			Email:           "wardanabayu456@gmail.com",
			Password:        utils.HashPasswordUser(passHash),
			EmailVerifiedAt: time.Now().Format("2006-01-02"),
			StatusID:        "2", // Misalnya, ID status user untuk Guru
		},
		{
			Name:            "Bayu Wardana 2",
			Username:        "9106012508950003",
			Email:           "wardanabayu457@gmail.com",
			Password:        utils.HashPasswordUser(passHash),
			EmailVerifiedAt: time.Now().Format("2006-01-02"),
			StatusID:        "3", // Misalnya, ID status user untuk Guru
		},
	}

	if err := db.Create(&guru).Error; err != nil {
		log.Fatal("Error creating guru:", err)
	}
	if err := db.Create(&user).Error; err != nil {
		log.Fatal("Error creating user:", err)
	}

	log.Println("Seeder Guru selesai")
}
