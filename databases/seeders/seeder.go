package seeders

import "gorm.io/gorm"

type Seeder interface {
	Run(db *gorm.DB)
}

func RunSeeders(db *gorm.DB) {
	GuruSeeder{}.Run(db)
	StatusUserSeeder{}.Run(db)
	// tambah seeder lain di sini
}
