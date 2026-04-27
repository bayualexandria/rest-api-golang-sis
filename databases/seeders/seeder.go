package seeders

import "gorm.io/gorm"

type Seeder interface {
	Run(db *gorm.DB)
}

func RunSeeders(db *gorm.DB) {
	StatusUserSeeder{}.Run(db)
	GuruSeeder{}.Run(db)
	SiswaSeeder{}.Run(db)
	// tambah seeder lain di sini
}
