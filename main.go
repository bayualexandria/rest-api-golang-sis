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
	// 1. Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("File .env tidak ditemukan: %v", err)
	}

	router := gin.Default()

	// ==========================================
	// 2. PASANG CORS DI POSISI PALING ATAS
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

	// 3. Koneksi ke database
	config.ConnectDatabase()

	// 4. Setup HTML Views (Opsional, jika menggunakan SSR/HTML Templates)
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

	if len(files) > 0 {
		router.LoadHTMLFiles(files...)
	}

	// ==========================================
	// 5. STATIC ROUTE & ROUTER LAINNYA
	// ==========================================
	router.Static("/storage", "./storage")

	// Setup routes web & API
	routes.SetupRouters(router)
	routes.SetupRoutersAPI(router)

	// 6. Seeders Handler
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		seeders.RunSeeders(config.DB)
		return // Keluar setelah seeding selesai
	}

	// 7. Services & Schedulers
	academicPeriodService := services.NewAcademicPeriodService(config.DB)
	scheduler.StartAcademicPeriodScheduler(academicPeriodService)

	absensiService := services.NewAbsensiService(config.DB)
	if err := absensiService.GenerateAlpa(); err != nil {
		log.Println("Generate ALPA gagal:", err)
	} else {
		log.Println("Generate ALPA berhasil dijalankan")
	}

	scheduler.StartAbsensiScheduler()

	
	// 9. Jalankan Server Gin
	// Gunakan variabel APP_PORT atau default ke ":8080"
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	}

	log.Printf("Server berjalan pada port %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
