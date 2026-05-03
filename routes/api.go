package routes

import (
	"backend-api/controllers"
	"backend-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutersAPI(app *gin.Engine) {
	// app.NoRoute(func(c *gin.Context) {
	// 	c.HTML(404, "404.html", gin.H{"message": "Halaman tidak ditemukan", "status": 404})
	// })
	route := app.Group("/api")
	{
		// Authentication Routes
		authRoute := route.Group("auth")
		authRoute.POST("/login-admin", controllers.LoginUserAdmin)
		authRoute.POST("/login", controllers.LoginUser)
		authRoute.GET("/verify/:email/:token", controllers.VerifyEmail)
		authRoute.POST("/forgot-password", controllers.ForgotPassword)
		authRoute.GET("/send-reset-password/:email/:token", controllers.SendResetPassword)
		authRoute.POST("/reset-password/:email/:token", controllers.ResetPassword)
		authRoute.POST("/logout", middleware.AuthMiddleware(), controllers.LogoutUser)

		// Drop Email Verified At
		route.GET("/drop-email-verified-at/:username", controllers.DropEmailVerifiedAt)

		// Login Social Media Routes
		route.GET("/login/google/:email/:idGoogle/:nameGoogle", controllers.LoginUserSocialMedia)

		// Endpoint Routes
		route.GET("/", controllers.HomeHandler)
		route.PATCH("/siswa/:nis", controllers.UpdateSiswa)
		route.GET("/user", middleware.AuthMiddleware(), controllers.GetUsers)
		route.GET("/user/:nis/siswa", middleware.AuthMiddleware(), middleware.RoleMiddleware(3, 4), controllers.GetUsersByNIS)
		route.GET("/siswa", middleware.AuthMiddleware(), middleware.AuthMiddleware(), controllers.GetSiswa)
		route.GET("/siswa/:nis", middleware.AuthMiddleware(), controllers.GetSiswaByID)
		route.GET("/guru", middleware.AuthMiddleware(), controllers.GetGuru)
		route.GET("/guru/:nip", middleware.AuthMiddleware(), controllers.GetGuruById)
		route.POST("/logout", middleware.AuthMiddleware(), controllers.LogoutUser)
	}

}
