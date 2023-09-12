package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"tp-core/middleware"
)

func InitRouter() {
	e := echo.New()
	e.Use(middleware.MiddlewareLogging)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "dunia tipu2",
		})
	})

	api := e.Group("/api/v1")

	api.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	api.GET("/callback/ds", func(c echo.Context) error {
		query := c.QueryParams()
		fmt.Println("query", query)
		return c.String(200, "success")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
