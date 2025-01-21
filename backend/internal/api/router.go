package api

import (
	"github.com/gin-gonic/gin"
	"github.com/Shiluco/UniTimetable/backend/internal/api/handler"
	"github.com/Shiluco/UniTimetable/backend/internal/api/middleware"
	"github.com/Shiluco/UniTimetable/backend/ent"
)

func SetupRoutes(r *gin.Engine, client *ent.Client) {
	// ミドルウェア設定
	r.Use(middleware.Logger())

	// ハンドラーの初期化
	userHandler := handler.NewUserHandler(client)

	// ユーザー関連のエンドポイント
	users := r.Group("/api/users")
	{
		users.GET("", userHandler.GetUsers)         // ユーザー一覧取得
		users.GET("/:id", userHandler.GetUser)      // 特定のユーザー取得
		users.POST("", userHandler.CreateUser)      // ユーザー作成
		users.PUT("/:id", userHandler.UpdateUser)   // ユーザー更新
		users.DELETE("/:id", userHandler.DeleteUser) // ユーザー削除
	}

	// その他のエンドポイント
	r.POST("/login", handler.Login)
	r.GET("/schedule", handler.GetSchedule)
}
