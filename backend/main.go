package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/Shiluco/UniTimetable/backend/internal/api"
	// "github.com/Shiluco/UniTimetable/backend/internal/api/middleware"
	"github.com/Shiluco/UniTimetable/backend/config"
	"github.com/Shiluco/UniTimetable/backend/pkg/logger"
	"github.com/Shiluco/UniTimetable/backend/ent"
	_ "github.com/lib/pq"
)

func main() {
	// 設定ファイルを読み込み
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("設定ファイルの読み込みに失敗しました: %v", err)
	}

	// ロガーの初期化
	logger.InitLogger()

	// entクライアントの初期化
	client, err := ent.Open("postgres", config.Config.DatabaseURL)
	if err != nil {
		log.Fatalf("entクライアントの作成に失敗しました: %v", err)
	}
	defer client.Close()

	// データベースのマイグレーション実行
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("スキーマの作成に失敗しました: %v", err)
	}

	// APIエンドポイントの設定
	r := api.SetupRoutes(client)

	// サーバーを起動
	port := config.Config.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("サーバーを起動します。ポート: %s", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
