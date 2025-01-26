package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

// InitLogger ロガーの初期化
func InitLogger() {
	// ログファイルの設定
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	// ロガーの設定
	Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo ログ出力
func LogInfo(message string) {
	Logger.Println("INFO: " + message)
}
