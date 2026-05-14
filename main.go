package main

import (
	"backend-api/config"
	"backend-api/databases/seeders"
	"backend-api/middleware"
	"backend-api/routes"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("File .env tidak ditemukan: %v", err)
	}
	router := gin.New()

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

	router.Use(middleware.CORSMiddleware())
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

	router.Run(os.Getenv("APP_URL"))
}
