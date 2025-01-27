package middleware

import (
	"errors"
	"fmt"
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
			// カスタムエラーを定義して c.Error でラップしておく
			err := errors.New("authorization header is missing or invalid prefix")
			c.Error(fmt.Errorf("AuthMiddleware error: %w", err))

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerSchema)
		claims, err := auth.ValidateToken(tokenString, "access_secret")
		if err != nil {
			// ValidateToken が失敗した場合
			c.Error(fmt.Errorf("ValidateToken error: %w", err))

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		user, err := client.User.Get(c.Request.Context(), claims.UserID)
		if err != nil {
			// DB からユーザーを取得できなかった場合
			c.Error(fmt.Errorf("User retrieval error: %w", err))

			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// 認証成功時はユーザー情報をセット
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
