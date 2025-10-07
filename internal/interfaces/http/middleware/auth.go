package middleware

import (
	"net/http"
	"strings"

	"taskqueue/internal/application/handlers"

	"github.com/gin-gonic/gin"
)

func JWTAuth(authApp *handlers.AuthApp) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "токен отсутствует"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := authApp.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "недействительный токен"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "роль пользователя не найдена"})
			c.Abort()
			return
		}

		if userRole != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error":    "недостаточно прав",
				"required": requiredRole,
				"current":  userRole,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func RequireSuper() gin.HandlerFunc {
	return RequireRole("super")
}
