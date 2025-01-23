package api

import (
    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/ent"
    "github.com/Shiluco/UniTimetable/backend/internal/api/handler"
	"github.com/Shiluco/UniTimetable/backend/internal/api/middleware"
    "github.com/Shiluco/UniTimetable/backend/internal/repository"
    "github.com/Shiluco/UniTimetable/backend/internal/domain/service"
)

func SetupRoutes(client *ent.Client) *gin.Engine {
    r := gin.New()

	r.Use(middleware.Logger())
    // 依存関係の注入
    userRepo := repository.NewUserRepository(client)
    userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
    userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

    // ルーティング
    api := r.Group("/api")
    {
        // 認証不要のエンドポイント
        auth := api.Group("/auth")
        {
            auth.POST("/login", authHandler.Login)
            auth.POST("/register", authHandler.Register)
        }

        // 認証が必要なエンドポイント
        users := api.Group("/users")
        users.Use(middleware.AuthMiddleware())
        {
            // コレクション操作
            users.GET("", userHandler.GetUsers)         // 一覧取得
            users.GET("/search", userHandler.SearchUsers) // 検索（クエリパラメータで検索タイプを指定）

            // 個別リソース操作
            users.GET("/:id", userHandler.GetUser)     // 個別取得
            users.PUT("/:id", userHandler.UpdateUser)  // 更新
            users.DELETE("/:id", userHandler.DeleteUser) // 削除
        }
    }

    return r
}
