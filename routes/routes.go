package routes

import (
	"backend-api/controllers"
	"backend-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouters(app *gin.Engine) {
	route := app.Group("/api")
	{
		route.GET("/", controllers.HomeHandler)
		route.POST("/login", controllers.LoginUser)
		route.GET("/user", middleware.AuthMiddleware(), controllers.GetUsers)
		route.GET("/siswa", middleware.AuthMiddleware(), controllers.GetSiswa)
		route.GET("/guru", middleware.AuthMiddleware(), controllers.GetGuru)
		route.GET("/guru/:nip", middleware.AuthMiddleware(), controllers.GetGuruById)
		route.POST("/logout", middleware.AuthMiddleware(), controllers.LogoutUser)
	}

}
