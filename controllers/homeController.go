package controllers

import (
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"data":    []string{},
		"message": "Data berhasil ditampilkan!",
		"success": true,
	})
}
