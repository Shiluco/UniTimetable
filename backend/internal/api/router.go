package api

import (
    "github.com/gin-gonic/gin"
    "github.com/Shiluco/UniTimetable/backend/ent"
    "github.com/Shiluco/UniTimetable/backend/internal/api/handler"
    "github.com/Shiluco/UniTimetable/backend/internal/api/middleware"
)

func SetupRoutes(client *ent.Client) *gin.Engine {
    // デフォルトのロガーを無効化
    gin.SetMode(gin.ReleaseMode)
    r := gin.New()

    // ミドルウェアの設定
    r.Use(gin.Recovery())  // パニックからの復帰
    r.Use(middleware.Logger())  // カスタムロガー
    r.Use(middleware.CORS())    // CORS

    // ハンドラーの初期化
    authHandler := handler.NewAuthHandler(client)
    scheduleHandler := handler.NewScheduleHandler(client)
    postHandler := handler.NewPostHandler(client)
	departmentHandler := handler.NewDepartmentHandler(client)
	majorHandler := handler.NewMajorHandler(client)
    userHandler := handler.NewUserHandler(client)
    //fileHandler := handler.NewFileHandler(client)

    // APIグループ
    api := r.Group("/api")
    {
        // 認証不要なエンドポイント
        auth := api.Group("/auth")
        {
            auth.POST("/register", authHandler.Register)
            auth.POST("/login", authHandler.Login)
        }
		// 学部関連（認証不要）
		departments := api.Group("/departments")
		{
    		departments.GET("", departmentHandler.GetDepartments)
    		departments.GET("/:id", departmentHandler.GetDepartment)
		}
		// 学科関連（認証不要）
		majors := api.Group("/majors")
		{
    		majors.GET("", majorHandler.GetMajors)
    		majors.GET("/:id", majorHandler.GetMajor)
		}
        // 認証が必要なエンドポイント
        authenticated := api.Group("")
        authenticated.Use(middleware.AuthMiddleware(client))
        {
            // ユーザー関連
            users := authenticated.Group("/users")
            {
                users.GET("/:id", userHandler.GetUsers)     // 個別取得（?total=trueで全件取得）
                users.PUT("/:id", userHandler.UpdateUser)   // 更新
                users.DELETE("/:id", userHandler.DeleteUser) // 削除
            }

            // 投稿関連
            posts := authenticated.Group("/posts")
            {
                posts.GET("", postHandler.GetPosts)         // 一覧取得
                posts.GET("/:id", postHandler.GetPosts)     // 単一投稿取得
                posts.POST("", postHandler.CreatePost)      // 投稿作成
                posts.PUT("/:id", postHandler.UpdatePost)   // 投稿更新
                posts.DELETE("/:id", postHandler.DeletePost) // 投稿削除
            }

            // 時間割関連
            schedules := authenticated.Group("/schedules")
            {
                schedules.GET("/:id", scheduleHandler.GetSchedule)
                schedules.POST("", scheduleHandler.CreateSchedule)
                schedules.PUT("/:id", scheduleHandler.UpdateSchedule)
                schedules.DELETE("/:id", scheduleHandler.DeleteSchedule)
            }

            // files := authenticated.Group("/files")
            // {
            //     files.POST("/upload", fileHandler.UploadFile)
            // }
        }
    }

    return r
}
