package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		msg := os.Getenv("MSG")
		return c.JSON(http.StatusOK, msg)
	})

	e.GET("/api-version", func(c echo.Context) error {
		msg := os.Getenv("API_VERSION")
		return c.JSON(http.StatusOK, msg)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
