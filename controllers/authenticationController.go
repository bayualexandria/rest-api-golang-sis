package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"backend-api/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username tidak sesuai"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password tidak sesuai"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat token"})
		return
	}
	var inputToken models.PersonalAccessToken
	inputToken.Token = token
	inputToken.TokenableType = "User"
	inputToken.TokenableID = user.ID
	inputToken.Name = "Personal Access Token"
	inputToken.Abilities = "*"
	inputToken.LastUsedAt = time.Now().Format("2006-01-02 15:04:05")
	inputToken.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	inputToken.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	config.DB.Create(&inputToken)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func LogoutUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var idToken models.PersonalAccessToken
	if err := config.DB.Where("token = ?", strings.TrimPrefix(token, "Bearer ")).First(&idToken).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	config.DB.Delete(&idToken)
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil logout"})
}
