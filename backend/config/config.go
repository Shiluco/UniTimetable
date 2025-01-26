package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config *AppConfig

type AppConfig struct {
	Port        string
	DatabaseURL string
}

func LoadConfig() error {
	// 開発環境の場合のみ.envを読み込む
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: .env file not found")
		}
	}

	// 環境変数から設定を読み込む
	Config = &AppConfig{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgresql://postgres:password@db:5432/unitimetable?sslmode=disable"),
	}

	return nil
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
