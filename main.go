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

	// ==========================================
	// 1. PASANG CORS DI POSISI PALING ATAS
	// ==========================================
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://sis.dev-coding.web.id",
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
			"Accept",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 2. Koneksi ke database
	config.ConnectDatabase()

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

	// ==========================================
	// 3. STATIC ROUTE & ROUTER LAINNYA
	// ==========================================
	router.Static("/storage", "./storage")

	// Setup routes web
	routes.SetupRouters(router)

	// Setup routes API
	routes.SetupRoutersAPI(router)

	// Seeders
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		seeders.RunSeeders(config.DB)
	}

	academicPeriodService := services.NewAcademicPeriodService(config.DB)
	scheduler.StartAcademicPeriodScheduler(
		academicPeriodService,
	)

	absensiService := services.NewAbsensiService(config.DB)
	if err := absensiService.GenerateAlpa(); err != nil {
		log.Println("Generate ALPA gagal:", err)
	} else {
		log.Println("Generate ALPA berhasil dijalankan")
	}

	scheduler.StartAbsensiScheduler()
	router.Run(os.Getenv("APP_URL"))
}
