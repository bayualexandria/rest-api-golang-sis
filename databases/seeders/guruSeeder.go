package seeders

import (
	"log"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

type GuruSeeder struct {
	gorm.Model
	gorm.DeletedAt
	Nip          string    `faker:"cc_number"`
	Nama         string `faker:"name"`
	JenisKelamin string `faker:"oneof: Laki-laki, Perempuan"`
	NoHp         string `faker:"phone_number"`
	ImageProfile string `faker:"url"`
	Alamat       string `faker:"mac_address"`
}
func (GuruSeeder) TableName() string {
	return "guru" // jadi singular
}

func SeederGuru(db *gorm.DB) {
	var gurus []GuruSeeder
	for i := 0; i < 10; i++ {
		guru := GuruSeeder{}
		err := faker.FakeData(&guru)
		if err != nil {
			log.Fatal(err)
		}
		gurus = append(gurus, guru)
	}
	db.Create(&gurus)
}
