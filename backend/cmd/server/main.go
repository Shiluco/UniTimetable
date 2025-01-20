package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/backend/internal/api/handler"
	"github.com/backend/internal/api/middleware"
	"github.com/backend/internal/config"
	"github.com/backend/pkg/logger"
)

func main() {
	// 設定ファイルを読み込み
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// ロガーの初期化
	logger.InitLogger()

	// Ginのエンジンを作成
	r := gin.Default()

	// ミドルウェアの設定
	r.Use(middleware.Logger())
	r.Use(middleware.Auth())

	// APIエンドポイントの設定
	handler.SetupRoutes(r)

	// サーバーを起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // デフォルトポート
	}

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
