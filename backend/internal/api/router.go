package api

import (
	"github.com/gin-gonic/gin"
	"github.com/backend/internal/api/handler"
	"github.com/backend/internal/api/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// ミドルウェア設定
	r.Use(middleware.Logger())

	// APIエンドポイント
	r.POST("/login", handler.Login)
	r.GET("/user/:id", handler.GetUser)
	r.POST("/post", handler.CreatePost)
	r.GET("/schedule", handler.GetSchedule)
}
