package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Age     int    `form:"age" validate:"required,gt=10,max=100"`
	Name    string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}

func main() {
	serve := gin.Default()

	serve.GET("/testing", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			//c.String(500, "%v",err)
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		//c.String(200, "%v", person)
		c.JSON(200, gin.H{"person": person})
	})

	serve.Run(":3000")
}
