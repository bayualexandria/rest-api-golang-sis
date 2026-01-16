package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouters(app *gin.Engine) {
	router := app.Group("")
	// Page not found
	app.NoRoute(func(c *gin.Context) {
		c.HTML(404,"404.html", gin.H{"message": "Halaman tidak ditemukan", "status": 404})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome to Backend API",
		})
	})
}
