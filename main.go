package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keisuke-333/go_practice/handler"
)

func main() {
	// インスタンス生成
	r := gin.Default()

	// プロジェクト内のhtmlを利用する
	r.LoadHTMLGlob("templates/*")

	// URLと処理の結び付け
	r.GET("/user/:id", handler.GetUser)

	r.Run(":3000")
}
