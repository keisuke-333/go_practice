package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
	ID   string
}

// ユーザーのリストを作って、リクエストで指定されたIDと一致するユーザーがいれば成功
func GetUser(c *gin.Context) {
	id := c.Param("id")

	userList := []User{
		{
			Name: "James",
			ID:   "1",
		},
	}

	for _, user := range userList {
		if id == user.ID {
			c.HTML(http.StatusOK, "sample.html", gin.H{
				"ID": user.ID,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user does not exist",
			})
		}
	}
}
