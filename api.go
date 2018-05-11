package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newAPIServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "password" {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/", func(c echo.Context) error {
		fmt.Println("got 'control.host/' request")
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
