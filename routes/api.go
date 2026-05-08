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

		// Users
		user := route.Group("/user")
		user.GET("/", middleware.AuthMiddleware(), controllers.GetUsers)
		user.GET("")
		user.GET("/guru/:nip", middleware.AuthMiddleware(), middleware.RoleMiddleware(2,3), controllers.GetGuruById)
		user.GET("/siswa/:nis", middleware.AuthMiddleware(), middleware.RoleMiddleware(4), controllers.GetUsersByNIS)

		// Siswa
		siswa := route.Group("/siswa")
		siswa.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.GetSiswa)
		siswa.GET("/:nis", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.GetSiswaByID)
		siswa.PATCH("/:nis", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.UpdateSiswa)

		// Guru
		guru := route.Group("/guru")
		guru.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.GetGuru)
		guru.PATCH("/:nip", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.UpdateGuru)

		// Logout
		route.POST("/logout", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.LogoutUser)
	}

}
