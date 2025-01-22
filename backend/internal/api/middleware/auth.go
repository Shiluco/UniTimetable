package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/internal/util"
)

// AuthMiddleware JWT認証ミドルウェア
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Authorizationヘッダーを取得
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
            c.Abort()
            return
        }

        // Bearer トークンの形式を確認
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
            c.Abort()
            return
        }

        // トークンを検証
        claims, err := util.ValidateJWT(parts[1])
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        // ユーザー情報をコンテキストに設定
        c.Set("username", claims["username"])
        c.Next()
    }
}
