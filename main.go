package main

import (
	"backend-api/config"
	"backend-api/routes"
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

	// 1. Koneksi ke database
	config.ConnectDatabase()
	config.EmailConfig()
	var files []string
	filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*", "http://my-server-dns", "http://192.168.88.103"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
		MaxAge:       12 * time.Hour,
	}))
	router.LoadHTMLFiles(files...)
	// Setup routes web
	routes.SetupRouters(router)

	// Setup routes API
	routes.SetupRoutersAPI(router)

	router.Run(os.Getenv("APP_URL"))
}
