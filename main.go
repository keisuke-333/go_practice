package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/user", show)

	e.Logger.Fatal(e.Start(":8080"))
}

func show(c echo.Context) error {
	users := []User{
		{
			Name:  "一郎",
			Email: "ichiro@test.com",
		},
		{
			Name:  "二郎",
			Email: "jiro@test.com",
		},
		{
			Name:  "三郎",
			Email: "saburo@test.com",
		},
	}
	return c.JSON(http.StatusOK, users)
}
