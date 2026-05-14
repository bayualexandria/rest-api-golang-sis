package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
        
        // Terima header apa pun yang diminta oleh Axios (penting untuk preflight)
        requestedHeaders := c.Request.Header.Get("Access-Control-Request-Headers")
        if requestedHeaders != "" {
            c.Writer.Header().Set("Access-Control-Allow-Headers", requestedHeaders)
        } else {
            c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
        }

        // Jika metodenya OPTIONS, berikan respon 204 (No Content) dan stop di sini
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}