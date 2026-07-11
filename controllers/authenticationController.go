package controllers

import (
	"backend-api/config"
	"backend-api/models"
	"backend-api/notifications"
	"backend-api/utils"
	"backend-api/validations"
	adminlogin "backend-api/validations/adminLogin"
	"backend-api/validations/siswalogin"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LoginUserAdmin(c *gin.Context) {
	var input adminlogin.LoginAdminValidation

	if err := c.ShouldBind(&input); err != nil {
		msg := adminlogin.TranslateErrorLoginAdmin(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": msg, "status": 401})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"message": "Username belum terdaftar", "status": 403})
		return
	}

	if user.StatusId != 1 && user.StatusId != 2 && user.StatusId != 3 {
		c.JSON(403, gin.H{"message": "User ini tidak memiliki akses login!", "status": 403})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(403, gin.H{"message": "Password yang anda masukan salah", "status": 403})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memuat token"})
		return
	}
	var inputToken models.PersonalAccessToken
	inputToken.Token = token
	inputToken.TokenableType = "User"
	inputToken.TokenableID = input.Username
	inputToken.Name = "Personal Access Token"
	inputToken.Abilities = "*"
	inputToken.LastUsedAt = time.Now().Format("2006-01-02 15:04:05")
	inputToken.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	inputToken.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err := config.DB.Table("users").Where("username = ?", input.Username).Where("email_verified_at", nil).First(&user).Error; err == nil || user.EmailVerifiedAt == "" {
		notifications.NotifikasiAktivasiAkunUser(user.Email, user.Name, "Silahkan verifikasi email anda untuk mengaktifkan akun anda, dengan cara klik link dibawah ini: ", os.Getenv("APP_URL")+"/api/auth/verify/"+user.Email+"/"+token)
		config.DB.Create(&inputToken)
		c.JSON(http.StatusOK, gin.H{"message": "Email belum terverifikasi, silakan cek email anda untuk verifikasi.", "status": 200})
		return
	}

	config.DB.Create(&inputToken)

	/// Di dalam controller saat LOGIN BERHASIL:
	// Gunakan properti cookie standar go untuk mengatur SameSite secara eksplisit
	c.Writer.Header().Set("Set-Cookie", "access_token="+token+"; Max-Age=86400; Path=/; SameSite=Lax; HttpOnly")

	// ATAU jika menggunakan c.SetCookie bawaan Gin, pastikan parameternya seperti ini:
	// c.SetCookie("access_token", token, 86400, "/", "", false, true)
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"message": "Username belum terdaftar", "status": 403})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": gin.H{"name": user.Name, "status_id": user.StatusId}, "message": "Anda berhasil login!", "status": 200}) // Hilangkan token dari body JSON
}

func LoginUser(c *gin.Context) {
	var input siswalogin.LoginSiswaValidation
	if err := c.ShouldBind(&input); err != nil {
		msg := siswalogin.TranslateErrorLoginSiswa(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": msg, "status": 401})
		return
	}
	var user models.User

	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"message": "Username belum terdaftar", "status": 403})
		return
	}

	if user.StatusId != 4 {
		c.JSON(403, gin.H{"message": "User ini tidak memiliki akses login!", "status": 403})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(403, gin.H{"message": "Password yang anda masukan salah", "status": 403})
		return
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memuat token", "status": 403})
		return
	}
	var inputToken models.PersonalAccessToken
	inputToken.Token = token
	inputToken.TokenableType = "User"
	inputToken.TokenableID = user.Username
	inputToken.Name = "Personal Access Token"
	inputToken.Abilities = "*"
	inputToken.LastUsedAt = time.Now().Format("2006-01-02 15:04:05")
	inputToken.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	inputToken.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err := config.DB.Where("username = ?", input.Username).Where("email_verified_at", nil).First(&user).Error; err == nil || user.EmailVerifiedAt == "" {
		inputToken.Token = token
		inputToken.TokenableType = "User"
		inputToken.TokenableID = user.Username
		inputToken.Name = "Personal Access Token"
		inputToken.Abilities = "*"
		inputToken.LastUsedAt = time.Now().Format("2006-01-02 15:04:05")
		inputToken.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		inputToken.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		notifications.NotifikasiAktivasiAkunUser(user.Email, user.Name, "Silahkan verifikasi email anda untuk mengaktifkan akun anda, dengan cara klik link dibawah ini: ", os.Getenv("APP_URL")+"/api/auth/verify/"+user.Email+"/"+token)
		config.DB.Create(&inputToken)
		c.JSON(http.StatusOK, gin.H{"message": "Email belum terverifikasi, silakan cek email anda untuk verifikasi.", "status": 200})
		return
	}

	config.DB.Create(&inputToken)

	// Di dalam controller saat LOGIN BERHASIL:
	// Gunakan properti cookie standar go untuk mengatur SameSite secara eksplisit
	c.Writer.Header().Set("Set-Cookie", "access_token="+token+"; Max-Age=86400; Path=/; SameSite=Lax; HttpOnly")

	// ATAU jika menggunakan c.SetCookie bawaan Gin, pastikan parameternya seperti ini:
	// c.SetCookie("access_token", token, 86400, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"user": gin.H{"name": user.Name, "status_id": user.StatusId}, "message": "Anda berhasil login!", "status": 200})
}

func LogoutUserAdmin(c *gin.Context) {
	// === MODIFIKASI LOGOUT DI SINI ===
	// Ambil token langsung dari Cookie, bukan Header lagi
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var idToken models.PersonalAccessToken
	if err := config.DB.Where("token = ?", tokenString).First(&idToken).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	config.DB.Where("tokenable_id = ?", idToken.TokenableID).Delete(&models.PersonalAccessToken{})

	// Hapus cookie di browser dengan mengatur MaxAge ke -1
	c.SetCookie("access_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil logout"})
}

func LogoutUserSiswa(c *gin.Context) {
	// === MODIFIKASI LOGOUT DI SINI ===
	// Ambil token langsung dari Cookie, bukan Header lagi
	nis := c.Param("nis")
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var idToken models.PersonalAccessToken
	if err := config.DB.Where("token = ?", tokenString).First(&idToken).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	config.DB.Where("tokenable_id = ?", nis).Delete(&models.PersonalAccessToken{})

	// Hapus cookie di browser dengan mengatur MaxAge ke -1
	c.SetCookie("access_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil logout"})
}

func VerifyEmail(c *gin.Context) {
	tokenString := c.Param("token")
	email := c.Param("email")

	var personalAccessToken models.PersonalAccessToken

	token, err := utils.VerifyJWT(tokenString)
	if err != nil || !token.Valid || config.DB.Where("token = ?", tokenString).First(&personalAccessToken).Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid atau sudah kedaluwarsa", "status": 401})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid", "status": 401})
		return
	}
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"message": "User tidak ditemukan", "status": 403})
		return
	}
	user.EmailVerifiedAt = time.Now().Format("2006-01-02 15:04:05")
	config.DB.Where("tokenable_id = ?", user.Username).Delete(&personalAccessToken)
	config.DB.Model(&user).Where("email = ?", email).Update("email_verified_at", user.EmailVerifiedAt)
	c.JSON(http.StatusOK, gin.H{"message": "Email berhasil diverifikasi", "status": 200})
}

func ForgotPassword(c *gin.Context) {
	var input validations.ForgotPasswordValidation
	if err := c.ShouldBind(&input); err != nil {
		msg := validations.TranslateForgotPasswordError(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": msg, "status": 401})
		return
	}
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"message": "Email belum terdaftar", "status": 403})
		return
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal memuat token", "status": 500})
		return
	}
	var passwordResetToken models.PasswordResetToken
	config.DB.Where("email = ?", user.Email).Delete(&models.PasswordResetToken{})

	passwordResetToken.Email = user.Email
	passwordResetToken.Token = token
	passwordResetToken.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	notifications.SendLinnkResetPassword(user.Email, user.Name, "Silahkan verifikasi email anda untuk mengaktifkan akun anda, dengan cara klik link dibawah ini: ", os.Getenv("APP_URL")+"/api/auth/send-reset-password/"+user.Email+"/"+token, "Reset Password")
	config.DB.Create(&passwordResetToken)
	c.JSON(http.StatusOK, gin.H{"message": "Link reset password telah dikirim ke email anda.", "status": 200})
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
		c.JSON(403, gin.H{"message": "Token tidak valid atau email tidak ditemukan", "status": 403})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"message": "User tidak ditemukan", "status": 403})
		return
	}
	newPassword := ResetPasswordNotRandomString(12)
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat hash password", "status": 500})
		return
	}
	notifications.NotificationResetPassword(user.Email, user.Name, "Silahkan menggunakan password dibawah ini untuk melakukan akses login ke portal!", newPassword, user.Username)
	config.DB.Where("email = ?", email).Delete(&models.PasswordResetToken{})

	config.DB.Model(&user).Where("email = ?", email).Update("password", hashedPassword)
	c.JSON(http.StatusOK, gin.H{"message": "Silahkan cek email untuk melihat perubahan password!", "status": 200})

}

// func ResetPassword(c *gin.Context) {
// 	var input validations.ResetPasswordValidation
// 	if err := c.ShouldBind(&input); err != nil {
// 		msg := validations.TranslateResetPasswordError(err)
// 		c.JSON(http.StatusUnauthorized, gin.H{"message": msg, "status": 401})
// 		return
// 	}
// 	tokenString := c.Param("token")
// 	email := c.Param("email")
// 	token, err := utils.VerifyJWT(tokenString)
// 	if err != nil || !token.Valid {
// 		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid atau sudah kedaluwarsa", "status": 401})
// 		return
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || claims["user_id"] == nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token tidak valid", "status": 401})
// 		return
// 	}
// 	var user models.User
// 	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
// 		c.JSON(403, gin.H{"message": "User tidak ditemukan", "status": 403})
// 		return
// 	}
// 	hashedPassword, err := utils.HashPassword(input.Password)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat hash password", "status": 500})
// 		return
// 	}
// 	user.Password = hashedPassword
// 	user.EmailVerifiedAt = time.Now().Format("2006-01-02 15:04:05")
// 	config.DB.Where("email = ?", email).Delete(&models.PasswordResetToken{})
// 	config.DB.Save(&user)
// 	c.JSON(http.StatusOK, gin.H{"message": "Password berhasil direset", "status": 200})
// }

func ChangePassword(c *gin.Context) {
	var username = c.Param("username")
	var input validations.ChangePasswordValidation
	var user models.User

	if err := c.ShouldBind(&input); err != nil {
		msg := validations.TranslateChangePasswordError(err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": msg, "status": 401})
		return
	}

	if err := config.DB.Where("username", username).First(&user).Error; err != nil {
		c.JSON(403, gin.H{"message": "User tidak ditemukan", "status": 403})
		return
	}
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal membuat hash password", "status": 500})
		return
	}
	user.Password = hashedPassword
	config.DB.Model(&user).Where("username = ?", username).Update("password", user.Password)
	c.JSON(http.StatusOK, gin.H{"message": "Password berhasil diubah", "status": 200})

}
