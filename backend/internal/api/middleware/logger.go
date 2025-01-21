package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger ミドルウェアはHTTPリクエストのログを記録します
func Logger() gin.HandlerFunc {
	// zapロガーの初期化
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return func(c *gin.Context) {
		// リクエスト開始時間
		start := time.Now()

		// リクエストパスの取得
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 後続のハンドラーを実行
		c.Next()

		// リクエスト終了時間と処理時間の計算
		end := time.Now()
		latency := end.Sub(start)

		// ステータスコードの取得
		status := c.Writer.Status()

		// クライアントIPの取得
		clientIP := c.ClientIP()

		// メソッドの取得
		method := c.Request.Method

		// エラーメッセージの取得（もしあれば）
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}

		// ログの記録
		logger.Info("incoming request",
			zap.String("path", path),
			zap.Int("status", status),
			zap.String("method", method),
			zap.Duration("latency", latency),
			zap.String("client_ip", clientIP),
			zap.String("error", errorMessage),
		)

		// 開発環境用のコンソール出力
		if gin.Mode() == gin.DebugMode {
			fmt.Printf("[GIN] %v | %3d | %13v | %15s | %-7s %s\n%s",
				end.Format("2006/01/02 - 15:04:05"),
				status,
				latency,
				clientIP,
				method,
				path,
				errorMessage,
			)
		}
	}
}
