package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keisuke-333/go_practice/src/controller"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.IndexGet)

	router.Run(":8080")
}
