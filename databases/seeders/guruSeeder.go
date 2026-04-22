package seeders

import (
	"log"

	"github.com/bxcodec/faker/v4"
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

type UserSeeder struct {
	ID           int
	Name         string
	Username     string
	Email        string
	Password     string
	StatusUserID int
}

func (GuruSeeder) TableName() string {
	return "guru" // jadi singular
}

func (UserSeeder) TableName() string {
	return "users" // jadi singular
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (s GuruSeeder) Run(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		username := faker.CCNumber()
		nama := faker.Name()
		guru := GuruSeeder{
			ID:           i + 1,
			Nip:          username, // Contoh NIP dengan 8 digit
			Nama:         nama,
			JenisKelamin: faker.Gender(),
			NoHp:         faker.Phonenumber(),
			Alamat:       faker.Word(),
			ImageProfile: faker.URL(),
		}
		passHash := "password123" // Contoh password default
		user := UserSeeder{
			ID:           i + 1,
			Name:         nama,
			Username:     username,
			Email:        faker.Email(),
			Password:     hashPassword(passHash),
			StatusUserID: 3, // Misalnya, ID status user untuk Guru
		}

		if err := db.Create(&guru).Error; err != nil {
			log.Println("Gagal insert:", err)
		}
		if err := db.Create(&user).Error; err != nil {
			log.Println("Gagal insert user:", err)
		}
	}

	log.Println("Seeder Guru selesai")
}
