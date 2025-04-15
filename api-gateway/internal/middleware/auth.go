package middleware

import (
	"api-gateway/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	tokenUtil utils.TokenUtil
}

func NewAuthMiddleware(tokenUtil utils.TokenUtil) *AuthMiddleware {
	return &AuthMiddleware{tokenUtil: tokenUtil}
}

func (m *AuthMiddleware) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
			return
		}

		claims, err := m.tokenUtil.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Add user ID to context for use in handlers
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
