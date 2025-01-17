package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのデフォルトのルーターを作成
	r := gin.Default()

	// "/" エンドポイントにハンドラーを登録
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// サーバーをデフォルトのポート8080で起動
	r.Run(":8080") // ポート番号を指定 (デフォルトで:8080)
}
