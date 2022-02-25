package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keisuke-333/go_practice/src/controller"
	"github.com/keisuke-333/go_practice/src/model"
)

func main() {
	r := gin.Default()

	// Connect to database
	model.DBConnect()

	// Routes
	r.GET("/tasks", controller.FindTasks)

	// Run the server
	r.Run(":8080")
}
