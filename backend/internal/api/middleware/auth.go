package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/ent"
    "github.com/Shiluco/UniTimetable/backend/internal/auth"
)

const (
    AuthorizationHeader = "Authorization"
    BearerSchema       = "Bearer "
    UserKey            = "user"
)

// AuthMiddleware 認証ミドルウェア
func AuthMiddleware(client *ent.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Authorizationヘッダーの取得
        authHeader := c.GetHeader(AuthorizationHeader)
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        // Bearer スキーマの確認
        if !strings.HasPrefix(authHeader, BearerSchema) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
            c.Abort()
            return
        }

        // トークンの取得
        tokenString := strings.TrimPrefix(authHeader, BearerSchema)

        // トークンの検証
        claims, err := auth.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // ユーザーの取得
        user, err := client.User.Get(c.Request.Context(), claims.UserID)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            c.Abort()
            return
        }

        // ユーザー情報をコンテキストに保存
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
