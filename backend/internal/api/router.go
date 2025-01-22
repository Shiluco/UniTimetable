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
    userHandler := handler.NewUserHandler(userService)

    // ルーティング
    api := r.Group("/api/v1")
    {
        users := api.Group("/users")
        {
            users.POST("", userHandler.CreateUser)
            users.GET("", userHandler.GetUsers)
            users.GET("/:id", userHandler.GetUser)
            users.PUT("/:id", userHandler.UpdateUser)
            users.DELETE("/:id", userHandler.DeleteUser)
        }
    }

    return r
}
