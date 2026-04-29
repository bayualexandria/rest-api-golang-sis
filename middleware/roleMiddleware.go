package middleware

import (
	"backend-api/models"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userData, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized!", "status": 401})
			return
		}

		user := userData.(models.User)

		// cek apakah role user ada di allowedRoles
		for _, role := range allowedRoles {
			if user.StatusUserId == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(403, gin.H{"message": "User ini tidak memiliki akses role!", "status": 403})
	}
}
