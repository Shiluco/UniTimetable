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
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}

	// 設定の読み込み
	Config = &AppConfig{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
	}

	return nil
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
