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

		// メソッドの取得
		method := c.Request.Method

		// 後続のハンドラーを実行
		c.Next()

		// リクエスト終了時間と処理時間の計算
		end := time.Now()
		latency := end.Sub(start)

		// ステータスコードの取得
		statusCode := c.Writer.Status()

		// クライアントIPの取得
		clientIP := c.ClientIP()

		// ユーザーエージェントの取得
		userAgent := c.Request.UserAgent()

		// エラーメッセージの取得（もしあれば）
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}

		// ログの記録
		if len(c.Errors) > 0 {
			logger.Error("HTTP Request",
				zap.String("path", path),
				zap.String("method", method),
				zap.Int("status", statusCode),
				zap.Duration("latency", latency),
				zap.String("ip", clientIP),
				zap.String("user-agent", userAgent),
				zap.String("errors", errorMessage),
			)
		} else {
			logger.Info("HTTP Request",
				zap.String("path", path),
				zap.String("method", method),
				zap.Int("status", statusCode),
				zap.Duration("latency", latency),
				zap.String("ip", clientIP),
				zap.String("user-agent", userAgent),
			)
		}

		// 開発環境用のコンソール出力
		if gin.Mode() == gin.DebugMode {
			fmt.Printf("[GIN] %v | %3d | %13v | %15s | %-7s %s\n%s",
				end.Format("2006/01/02 - 15:04:05"),
				statusCode,
				latency,
				clientIP,
				method,
				path,
				errorMessage,
			)
		}
	}
}
