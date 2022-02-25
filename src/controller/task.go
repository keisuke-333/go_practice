package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keisuke-333/go_practice/src/model"
)

func FindTasks(c *gin.Context) {
	var tasks []model.Task
	model.DB.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}
