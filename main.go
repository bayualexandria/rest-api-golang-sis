package main

import (
	"backend-api/config"
	"backend-api/routes"
	"log"
	"os"

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

	routes.SetupRouters(router)
	router.Run(os.Getenv("APP_URL"))
}
