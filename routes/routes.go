package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouters(app *gin.Engine) {
	router := app.Group("")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome to Backend API",
		})
	})
}
