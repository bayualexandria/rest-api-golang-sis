package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	sqlite "gorm.io/driver/sqlite" // Tetap di-import, tapi kita override dengan modernc
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	db := os.Getenv("DB_CONNECTION") // ambil tipe database dari .env

	if db == "mysql" {
		database, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)), &gorm.Config{})
		if err != nil {
			log.Fatal("Gagal konek ke MySQL:", err)
		}
		fmt.Println("✅ Koneksi ke MySQL berhasil!")
		DB = database
		return
	}
	if db == "sqlite" {
		database, err := gorm.Open(sqlite.Dialector{
			DriverName: "sqlite", // modernc.org/sqlite pakai nama driver ini
			DSN:        "./databases/database.sqlite",
		}, &gorm.Config{})
		if err != nil {
			log.Fatal("Gagal konek ke SQLite:", err)
		}

		fmt.Println("✅ Koneksi ke SQLite (tanpa CGo) berhasil!")
		DB = database
		return
	}

	fmt.Println("✅ Tidak ada koneksi ke database!")

}
