package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// CORS CORSミドルウェアを設定
func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",     // Next.js開発サーバー
		"http://localhost:6006",     // Storybook
	}
	config.AllowMethods = []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Content-Type",
		"Accept",
		"Authorization",
	}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{
		"Content-Length",
	}
	config.MaxAge = 86400 // プリフライトリクエストのキャッシュ時間（24時間）

	return cors.New(config)
} 