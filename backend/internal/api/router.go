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
            users.GET("", userHandler.GetUsers)
            users.GET("/:id", userHandler.GetUser)
            users.PUT("/:id", userHandler.UpdateUser)
            users.DELETE("/:id", userHandler.DeleteUser)
            users.GET("/by-email", userHandler.GetUserByEmail)
            users.GET("/by-name", userHandler.GetUserByName)
            users.GET("/search/name", userHandler.SearchUsersByName)
            users.GET("/search/email", userHandler.SearchUsersByEmail)
        }
    }

    return r
}
