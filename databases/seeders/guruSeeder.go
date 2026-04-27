package seeders

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GuruSeeder struct {
	ID           int
	Nip          string
	Nama         string
	JenisKelamin string
	NoHp         string
	Alamat       string
	ImageProfile string
}

type UserGuruSeeder struct {
	ID              int
	Name            string
	Username        string
	Email           string
	Password        string
	EmailVerifiedAt string
	StatusUserID    int
}

func (GuruSeeder) TableName() string {
	return "guru" // jadi singular
}

func (UserGuruSeeder) TableName() string {
	return "users" // jadi singular
}

func hashPasswordUserGuru(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (s GuruSeeder) Run(db *gorm.DB) {

	guru := []GuruSeeder{
		{
			ID:           1,
			Nip:          "9106012508950001", // Contoh NIP dengan 8 digit
			Nama:         "Bayu Wardana",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567890",
			Alamat:       "Jl. Contoh Alamat No. 123, Kota Contoh",
			ImageProfile: "https://example.com/images/bayu.jpg",
		},
		{
			ID:           2,
			Nip:          "9106012508950002", // Contoh NIP dengan 8 digit
			Nama:         "Bayu Wardana",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567891",
			Alamat:       "Jl. Contoh Alamat No. 123, Kota Contoh",
			ImageProfile: "https://example.com/images/bayu.jpg",
		},
		{
			ID:           3,
			Nip:          "9106012508950003", // Contoh NIP dengan 8 digit
			Nama:         "Bayu Wardana",
			JenisKelamin: "Laki-laki",
			NoHp:         "081234567892",
			Alamat:       "Jl. Contoh Alamat No. 123, Kota Contoh",
			ImageProfile: "https://example.com/images/bayu.jpg",
		},
	}
	passHash := "admin123" // Contoh password default
	user := []UserGuruSeeder{
		{
			ID:              1,
			Name:            "Bayu Wardana",
			Username:        "9106012508950001",
			Email:           "wardanabayu455@gmail.com",
			Password:        hashPasswordUserGuru(passHash),
			EmailVerifiedAt: time.Now().Format("2006-01-02"),
			StatusUserID:    1, // Misalnya, ID status user untuk Guru
		},
		{
			ID:              2,
			Name:            "Bayu Wardana",
			Username:        "9106012508950002",
			Email:           "wardanabayu456@gmail.com",
			Password:        hashPasswordUserGuru(passHash),
			EmailVerifiedAt: time.Now().Format("2006-01-02"),
			StatusUserID:    2, // Misalnya, ID status user untuk Guru
		},
		{
			ID:              3,
			Name:            "Bayu Wardana",
			Username:        "9106012508950003",
			Email:           "wardanabayu457@gmail.com",
			Password:        hashPasswordUserGuru(passHash),
			EmailVerifiedAt: time.Now().Format("2006-01-02"),
			StatusUserID:    3, // Misalnya, ID status user untuk Guru
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
