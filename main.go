package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keisuke-333/go_practice/article"
	"github.com/keisuke-333/go_practice/handler"
)

func main() {
	article := article.New()
	r := gin.Default()
	r.GET("/article", handler.ArticlesGet(article))
	r.POST("/article", handler.ArticlePost(article))

	r.Run()

}
