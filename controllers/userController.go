package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUsersByNIS(c *gin.Context) {
	nis := c.Param("nis")
	var user models.User
	result := config.DB.Where("nis = ?", nis).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
