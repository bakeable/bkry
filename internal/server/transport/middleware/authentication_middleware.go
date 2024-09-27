package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthenticatedRoutes() gin.HandlerFunc {
    return func(c *gin.Context) {
        // authHeader := c.GetHeader("Authorization")
        // idToken := strings.TrimPrefix(authHeader, "Bearer ")

        // fbUser, err := firebase.GetAuthClient().VerifyIDToken(context.Background(), idToken)
        // if err != nil {
        //     c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing token"})
        //     return
        // }

        // c.Set("fbUser", fbUser)
		// c.Set("userId", fbUser.UID)
        c.Next()
    }
}