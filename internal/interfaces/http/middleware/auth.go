package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {

		userRole := c.GetHeader("X-User-Role")

		c.Header("Debug-Role", userRole)

		if userRole == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
				"debug": "X-User-Role header is empty",
			})
			c.Abort()
			return
		}

		if userRole != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error":    "insufficient permissions",
				"required": requiredRole,
				"current":  userRole,
			})
			c.Abort()
			return
		}

		c.Set("user_role", userRole)
		c.Next()
	}
}

func RequireSuper() gin.HandlerFunc {
	return RequireRole("super")
}
