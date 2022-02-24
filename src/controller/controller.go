package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexGet(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
