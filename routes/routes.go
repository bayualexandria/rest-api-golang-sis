package routes

import (
	"backend-api/controllers"
	"backend-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouters(app *gin.Engine) {
	route := app.Group("/api")
	{
		// Authentication Routes
		authRoute := route.Group("auth")
		authRoute.POST("/login-admin", controllers.LoginUserAdmin)
		authRoute.GET("/verify/:email/:token", controllers.VerifyEmail)

		// Drop Email Verified At
		route.GET("/drop-email-verified-at/:username", controllers.DropEmailVerifiedAt)

		route.GET("/", controllers.HomeHandler)
		authRoute.POST("/logout", middleware.AuthMiddleware(), controllers.LogoutUser)
		route.GET("/user", middleware.AuthMiddleware(), controllers.GetUsers)
		route.GET("/siswa", middleware.AuthMiddleware(), controllers.GetSiswa)
		route.GET("/guru", middleware.AuthMiddleware(), controllers.GetGuru)
		route.GET("/guru/:nip", middleware.AuthMiddleware(), controllers.GetGuruById)
		route.POST("/logout", middleware.AuthMiddleware(), controllers.LogoutUser)
	}

}
