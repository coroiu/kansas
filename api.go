package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func NewApiServer() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println("got 'control.host/' request")
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
