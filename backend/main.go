package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Shiluco/UniTimetable/backend/ent"
	"github.com/Shiluco/UniTimetable/backend/internal/api"
	"github.com/Shiluco/UniTimetable/backend/pkg/logger"
	_ "github.com/lib/pq"
)

func main() {
	logger.InitLogger()

	var client *ent.Client
	var err error
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		client, err = ent.Open("postgres", "postgresql://postgres:password@db:5432/unitimetable?sslmode=disable")
		if err == nil {
			break
		}
		log.Printf("DB接続再試行 (%d/%d): %v", i+1, maxRetries, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("スキーマ作成失敗: %v", err)
	}

	r := api.SetupRoutes(client)
	port := "8080"
	log.Printf("サーバー起動 ポート: %s", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("サーバー起動失敗: %v", err)
	}
	SaveUniversityData(client, "backend/internal/univ_data/univ.json")
}
