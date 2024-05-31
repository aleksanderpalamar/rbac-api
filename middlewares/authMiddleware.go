package middlewares

import (
	"net/http"

	"github.com/aleksanderpalamar/rbac-api/auth"
	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		claims, err := auth.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
