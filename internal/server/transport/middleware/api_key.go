package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// apiKey is the API key for the webhook
var apiKey string = "9yt84hoef42nFDsldadfjais32LlK213lads23"


func ApiKeyAuthentication() gin.HandlerFunc {
    return func(c *gin.Context) {
		// Check if the provided API key is correct
		if c.GetHeader("X-API-KEY") != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}
        c.Next()
    }
}