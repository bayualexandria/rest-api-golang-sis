package main

import (
	"backend-api/config"
	"backend-api/databases/seeders"
	"backend-api/routes"
	"backend-api/services"
	"backend-api/services/scheduler"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("File .env tidak ditemukan: %v", err)
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			// Tambahkan Domain running front-end jika API ini tidak bisa diakses oleh Front-End
			"http://localhost",
			"http://localhost:3000",
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// 1. Koneksi ke database
	config.ConnectDatabase()
	// config.EmailConfig()

	var files []string
	filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{"message": "Halaman tidak ditemukan", "status": 404})
	})
	router.LoadHTMLFiles(files...)
	// Setup routes web
	routes.SetupRouters(router)

	// Setup routes API
	routes.SetupRoutersAPI(router)

	// router.Use(middleware.CORSMiddleware())
	// Logger dan Recovery tetap diperlukan agar tidak crash

	router.Static("/storage", "./storage")

	// Seeders
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		seeders.RunSeeders(config.DB)
	}
	// Mailer
	// m := mailers.NewMailer()

	// m.Send(
	// 	"wardanabayu455@gmail.com",
	// 	"Test Email",
	// 	"<h1>Halo dari Golang</h1>",
	// )

	tahunAjaranService :=
		services.NewTahunAjaranService(config.DB)

	tahunAjaran, err :=
		tahunAjaranService.EnsureCurrentYear()

	if err != nil {
		log.Fatal(err)
	}
	semesterService :=
		services.NewSemesterService(config.DB)
	semesterService.EnsureCurrentSemester(tahunAjaran)

	absensiService := services.NewAbsensiService(config.DB)
	if err := absensiService.GenerateAlpa(); err != nil {
		log.Println("Generate ALPA gagal:", err)
	} else {
		log.Println("Generate ALPA berhasil dijalankan")
	}

	scheduler.StartAbsensiScheduler()
	router.Run(os.Getenv("APP_URL"))
}
