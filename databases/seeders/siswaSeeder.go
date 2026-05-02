package seeders

import (
	"log"
	"time"

	"github.com/bxcodec/faker/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SiswaSeeder struct {
	ID           int
	Nis          string
	Nama         string
	JenisKelamin string
	NoHp         string
	Alamat       string
	ImageProfile string
}

type UserSiswaSeeder struct {
	ID              int
	Name            string
	Username        string
	Email           string
	Password        string
	EmailVerifiedAt string
	StatusUserID    int
}

func (SiswaSeeder) TableName() string {
	return "siswa" // jadi singular
}

func (UserSiswaSeeder) TableName() string {
	return "users" // jadi singular
}

func HashPasswordUserSiswa(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (s SiswaSeeder) Run(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		genders := []string{"Laki-laki", "Perempuan"}
		username := faker.CCNumber()
		nama := faker.Name()
		siswa := SiswaSeeder{
			ID:           i + 1,
			Nis:          username, // Contoh NIS dengan 8 digit
			Nama:         nama,
			JenisKelamin: genders[i%2], // Alternatif jenis kelamin
			NoHp:         faker.Phonenumber(),
			Alamat:       faker.Word() + " Street No., " + faker.CCNumber(),
			ImageProfile: "storages/logo-pendidikan.png",
		}
		passHash := "admin123" // Contoh password default
		user := UserSiswaSeeder{
			ID:              i + 4,
			Name:            nama,
			Username:        username,
			Email:           faker.Email(),
			Password:        HashPasswordUserSiswa(passHash),
			EmailVerifiedAt: time.Now().Format("2006-01-02"),
			StatusUserID:    4, // Misalnya, ID status user untuk Siswa
		}
		if err := db.Create(&siswa).Error; err != nil {
			log.Fatal("Error creating siswa:", err)
		}
		if err := db.Create(&user).Error; err != nil {
			log.Fatal("Error creating user:", err)
		}
	}
	log.Println("Seeder Siswa selesai")
}
