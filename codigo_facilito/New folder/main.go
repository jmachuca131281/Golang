package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"net/http"
)

func main() {
	fmt.Println("Hola mundo")
	e := echo.New()
	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "hello, World")
	})
	e.Logger.Fatal(e.Start(":1323"))
}