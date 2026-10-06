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

		// Login Social Media Routes
		route.GET("/login-admin/google/:email/:idGoogle/:nameGoogle", controllers.LoginUserAdminSocialMedia)
		route.GET("/login/google/:email/:idGoogle/:nameGoogle", controllers.LoginUserSiswaSocialMedia)

		// Endpoint Routes
		route.GET("/", controllers.HomeHandler)

		// Profile Sekolah
		route.GET("/profile-sekolah", controllers.ProfileSekolahHandler)
		// Absensi Siswa

		route.POST("/absensi", controllers.AddAbsensiSiswa)

		// Personal Access Token
		route.DELETE("/access-token/:username", controllers.GetAccessToken)

		// =========================================================
		// ROUTES WITH MIDDLEWARE
		// =========================================================

		// Profile Sekolah
		route.PATCH("/profile-sekolah", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.UpdateProfileSekolahHandler)

		// Users
		user := route.Group("/user")
		user.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.GetUsers)
		user.GET("/:username", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.GetUsersByUsername)
		user.GET("/:username/guru", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetUsersByNIP)
		user.GET("/:username/siswa", middleware.AuthMiddleware(), middleware.RoleMiddleware(4), controllers.GetUsersByNIS)
		user.PUT("/change-password/:username", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.ChangePassword)
		user.GET("/:username/update-email-verified", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.UpdateEmailVerified)
		user.GET("/:username/update-email-verified-at", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.UpdateEmailVerifiedAt)

		// Siswa
		siswa := route.Group("/siswa")
		siswa.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetSiswa)
		siswa.GET("/search", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetSiswaSearch)
		siswa.GET("/:username", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3, 4), controllers.GetDataByNIS)
		siswa.POST("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.AddSiswa)
		siswa.PATCH("/:nis", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 4), controllers.UpdateSiswa)
		siswa.DELETE("/:nis", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.DeleteSiswa)

		// Siswa Kelas
		siswaKelas := route.Group("/siswa-kelas")
		siswaKelas.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetSiswaKelas)
		siswaKelas.GET("/:nis/:kelas", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetSiswaKelasByNis)
		siswaKelas.POST("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2), controllers.AddSiswaKelas)
		siswaKelas.DELETE("/:id/:nis", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2), controllers.DeleteSiswaKelas)

		// Ruang Kelas
		ruangKelas := route.Group("/ruang-kelas")
		ruangKelas.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.RuangKelas)
		ruangKelas.GET("/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2), controllers.RuangKelasById)
		ruangKelas.GET("/guru/:nip", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.RuangKelasByNip)
		ruangKelas.POST("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.AddRuangKelas)
		ruangKelas.PATCH("/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.UpdateRuangKelas)

		
		// Mata Pelajaran
		mapel := route.Group("/mapel")
		mapel.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.GetDataAllMapel)
		mapel.GET("/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.GetDataMapelById)
		mapel.POST("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.InsertDataMapel)
		mapel.PATCH("/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.UpdateDataMapel)

		// Guru Mapel
		guruMapel := route.Group("/guru-mapel")
		guruMapel.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetAllGuruMapel)
		guruMapel.POST("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.AddGuruMapel)

		// kelas
		kelas := route.Group("/kelas")
		kelas.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetKelas)

		// Guru
		guru := route.Group("/guru")
		guru.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetGuru)
		guru.GET("/:username", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetDataGuruByNIP)
		guru.POST("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.AddGuru)
		guru.PATCH("/:nip", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.UpdateGuru)
		guru.DELETE("/:nip", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.DeleteGuru)

		// Semester
		semester := route.Group("/semester")
		semester.GET("/", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.GetSemester)

		// Trash Data
		trash := route.Group("/trash")
		trash.GET("/siswa", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.GetTrashSiswa)
		trash.GET("/siswa/restore-all", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.RestoreDataTrashAllSiswa)
		trash.PATCH("/siswa/restore/:nis", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.RestoreDataTrashSiswa)
		trash.GET("/guru", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.GetTrashGuru)
		trash.GET("/guru/restore-all", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.RestoreDataTrashAllGuru)
		trash.PATCH("/guru/restore/:nip", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.RestoreDataTrashGuru)

		// Logout
		route.POST("/logout/:nis", middleware.AuthMiddleware(), middleware.RoleMiddleware(4), controllers.LogoutUserSiswa)
		route.POST("/logout-admin", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2, 3), controllers.LogoutUserAdmin)
	}
}
