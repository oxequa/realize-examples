package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

func main() {
	fmt.Println("Server listening on 0.0.0.0:3000")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World, Echo Framework!")
	})
	e.Run(standard.New("0.0.0.0:3000"))
	for {
		fmt.Println("test")
	}
}
