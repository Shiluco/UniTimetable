package middleware

import (
	"net/http"
	"strings"

	"github.com/Shiluco/UniTimetable/backend/ent"
	"github.com/Shiluco/UniTimetable/backend/internal/auth"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	BearerSchema        = "Bearer "
	UserKey             = "user"
)

// AuthMiddleware 認証ミドルウェア
func AuthMiddleware(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" || !strings.HasPrefix(authHeader, BearerSchema) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerSchema)
		claims, err := auth.ValidateToken(tokenString, "access_secret")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		user, err := client.User.Get(c.Request.Context(), claims.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set(UserKey, user)
		c.Next()
	}
}

// GetCurrentUser コンテキストから現在のユーザーを取得
func GetCurrentUser(c *gin.Context) *ent.User {
	user, exists := c.Get(UserKey)
	if !exists {
		return nil
	}
	return user.(*ent.User)
}
