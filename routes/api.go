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
	app.OPTIONS("/*any", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.AbortWithStatus(204)
	})
	route := app.Group("/api")
	{
		route.Use(middleware.CORSMiddleware())

		{
			// Authentication Routes
			authRoute := route.Group("auth")
			authRoute.POST("/login-admin", controllers.LoginUserAdmin)
			authRoute.POST("/login", controllers.LoginUser)
			authRoute.GET("/verify/:email/:token", controllers.VerifyEmail)
			authRoute.POST("/forgot-password", controllers.ForgotPassword)
			authRoute.GET("/send-reset-password/:email/:token", controllers.SendResetPassword)

			// Login Social Media Routes
			route.GET("/login-admin/google/:email/:idGoogle/:nameGoogle", controllers.LoginUserSocialMedia)

			// Endpoint Routes
			route.GET("/", controllers.HomeHandler)

			// Users
			user := route.Group("/user")
			user.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.GetUsers)
			user.GET("/:username", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.GetUsersByUsername)
			user.GET("/:username/guru", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetUsersByNIP)
			user.GET("/:username/siswa", middleware.AuthMiddleware(), middleware.RoleMiddleware(4), controllers.GetUsersByNIS)
			user.PUT("/change-password/:username", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.ChangePassword)

			// Siswa
			siswa := route.Group("/siswa")
			siswa.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetSiswa)
			siswa.GET("/:nis", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.GetSiswaByID)
			siswa.POST("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.AddSiswa)
			siswa.PATCH("/:nis", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.UpdateSiswa)

			// Guru
			guru := route.Group("/guru")
			guru.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetGuru)
			guru.PATCH("/:nip", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.UpdateGuru)

			// Logout
			route.POST("/logout", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.LogoutUser)
		}
	}

}
