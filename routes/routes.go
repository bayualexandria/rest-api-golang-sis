package routes

import (
	"backend-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouters(app *gin.Engine) {
	app.GET("/", controllers.HomeHandler)
	app.GET("/api/user", controllers.GetUsers)
	app.GET("/api/siswa", controllers.GetSiswa)
	app.GET("/api/guru", controllers.GetGuru)
	app.GET("/api/guru/:nip", controllers.GetGuruById)

}
