package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"backend-api/notifications"
	"backend-api/utils"
	"backend-api/validations"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LoginUserAdmin(c *gin.Context) {
	var input validations.LoginValidation

	if err := c.ShouldBindJSON(&input); err != nil {
		msg := validations.TranslateError(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"error": "Username belum terdaftar"})
		return
	}

	if config.DB.Where("status_id != ?", "4").First(&user).Error != nil {
		c.JSON(403, gin.H{"error": "User ini tidak memiliki akses login!"})
		return

	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(403, gin.H{"error": "Password yang anda masukan salah"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat token"})
		return
	}
	if err := config.DB.Where("username = ?", input.Username).Where("email_verified_at", nil).First(&user).Error; err == nil || user.EmailVerifiedAt == "" {
		notifications.SendVerificationEmail(user.Email, map[string]interface{}{
			"Name":             user.Name,
			"VerificationLink": "http://" + os.Getenv("APP_URL") + "/api/auth/verify/" + user.Email + "/" + token,
		})
		c.JSON(http.StatusOK, gin.H{"data": user.EmailVerifiedAt, "message": "Email belum terverifikasi, silakan cek email anda untuk verifikasi."})
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

	c.JSON(http.StatusOK, gin.H{"data": user, "token": token})
}

func LoginUser(c *gin.Context) {
	var input validations.LoginValidation
	if err := c.ShouldBindJSON(&input); err != nil {
		msg := validations.TranslateError(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}
	var user models.User

	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"error": "Username belum terdaftar"})
		return
	}

	if err := config.DB.Where("status_id = ?", "4").First(&user).Error; err != nil {
		c.JSON(403, gin.H{"error": "User ini tidak memiliki akses login!"})
		return

	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(403, gin.H{"error": "Password yang anda masukan salah"})
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
	c.JSON(http.StatusOK, gin.H{"data": user, "token": token})
}

func LogoutUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var idToken models.PersonalAccessToken
	if err := config.DB.Where("token = ?", strings.TrimPrefix(token, "Bearer ")).First(&idToken).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	config.DB.Where("tokenable_id = ?", idToken.TokenableID).Delete(&models.PersonalAccessToken{})
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil logout"})
}

func DropEmailVerifiedAt(c *gin.Context) {

	var user models.User
	if err := config.DB.Where("username = ?", c.Param("username")).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"error": "User tidak ditemukan"})
		return
	}
	user.EmailVerifiedAt = ""
	config.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Email verified at field has been cleared"})
}

func VerifyEmail(c *gin.Context) {
	tokenString := c.Param("token")
	email := c.Param("email")

	token, err := utils.VerifyJWT(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau sudah kedaluwarsa"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"error": "User tidak ditemukan"})
		return
	}
	user.EmailVerifiedAt = time.Now().Format("2006-01-02 15:04:05")
	config.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Email berhasil diverifikasi"})
}

func ForgotPassword(c *gin.Context) {
	var input validations.ForgotPasswordValidation
	if err := c.ShouldBindJSON(&input); err != nil {
		msg := validations.TranslateForgotPasswordError(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"error": "Email belum terdaftar"})
		return
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memuat token"})
		return
	}
	var passwordResetToken models.PasswordResetToken
	config.DB.Where("email = ?", user.Email).Delete(&models.PasswordResetToken{})
	notifications.SendResetPassword(user.Email, map[string]interface{}{
		"Name":              user.Name,
		"ResetPasswordLink": "http://" + os.Getenv("APP_URL") + "/api/auth/send-reset-password/" + user.Email + "/" + token,
	})
	passwordResetToken.Email = user.Email
	passwordResetToken.Token = token
	passwordResetToken.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	config.DB.Create(&passwordResetToken)
	c.JSON(http.StatusOK, gin.H{"message": "Link reset password telah dikirim ke email anda."})
}

func ResetPasswordNotRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[utils.RandomInt(0, len(letters))]
	}
	return string(result)
}

func SendResetPassword(c *gin.Context) {
	var tokenString = c.Param("token")
	var email = c.Param("email")
	var passwordResetToken models.PasswordResetToken
	if err := config.DB.Where("email = ? AND token = ?", email, tokenString).First(&passwordResetToken).Error; err != nil {
		c.JSON(403, gin.H{"error": "Token tidak valid atau email tidak ditemukan"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"error": "User tidak ditemukan"})
		return
	}
	newPassword := ResetPasswordNotRandomString(12)
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat hash password"})
		return
	}

	var data = map[string]interface{}{
		"Name":              user.Name,
		"Password":          newPassword,
		"ResetPasswordLink": "http://" + os.Getenv("APP_URL") + "/auth/reset-password/" + email + "/" + tokenString,
	}

	notifications.SendNewPassword(email, data)
	user.Password = hashedPassword
	config.DB.Save(&user)
	config.DB.Where("email = ?", email).Delete(&models.PasswordResetToken{})
	c.JSON(http.StatusOK, gin.H{"message": "Link reset password telah dikirim ke email anda."})

}

func ResetPasswordNotification(c *gin.Context) {

	var token = c.Param("token")
	var email = c.Param("email")

	var data = map[string]interface{}{
		"Name":              "User",
		"ResetPasswordLink": "http://" + os.Getenv("APP_URL") + "/auth/reset-password/" + email + "/" + token,
	}
	notifications.SendResetPassword(email, data)
	c.JSON(http.StatusOK, gin.H{"message": "Link reset password telah dikirim ke email anda."})
}

func ResetPassword(c *gin.Context) {
	var input validations.ResetPasswordValidation
	if err := c.ShouldBindJSON(&input); err != nil {
		msg := validations.TranslateResetPasswordError(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}
	tokenString := c.Param("token")
	email := c.Param("email")
	token, err := utils.VerifyJWT(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau sudah kedaluwarsa"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"error": "User tidak ditemukan"})
		return
	}
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat hash password"})
		return
	}
	user.Password = hashedPassword
	user.EmailVerifiedAt = time.Now().Format("2006-01-02 15:04:05")
	config.DB.Where("email = ?", email).Delete(&models.PasswordResetToken{})
	config.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Password berhasil direset"})
}
