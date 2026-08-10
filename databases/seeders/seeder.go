package seeders

import (
	"log"

	"gorm.io/gorm"
)

type Seeder interface {
	Run(db *gorm.DB)
}

func RunSeeders(db *gorm.DB) {
	if db == nil {
		log.Fatal("Database belum diinisialisasi!")
	}
	StatusSiswaSeeder{}.Run(db)
	StatusUserSeeder{}.Run(db)
	GuruSeeder{}.Run(db)
	SiswaSeeder{}.Run(db)
	KelasSeeder{}.Run(db)
	ProfileSekolahSeeder{}.Run(db)
	StatusKehadiranSeeder{}.Run(db)
	// tambah seeder lain di sini
}
