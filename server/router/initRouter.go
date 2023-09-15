package router

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"os"
	"tp-core/server/controllers"
	"tp-core/server/helpers"
	"tp-core/server/middleware"
	"tp-core/web"
)

func InitRouter() {
	e := echo.New()
	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler
	//field validation https://godoc.org/github.com/go-playground/validator
	custValidator := &helpers.CustomValidator{Validator: validator.New()}
	err := custValidator.Init()
	if err != nil {
		fmt.Printf("Failed to init validator: %v\n", err)
		os.Exit(1)
	}
	e.Validator = custValidator

	web.RegisterHandlers(e)

	api := e.Group("/api/v1")

	api.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	api.GET("/callback/ds", func(c echo.Context) error {
		query := c.QueryParams()
		fmt.Println("query", query)
		return c.String(200, "success")
	})

	api.GET("/callback/wpy", func(c echo.Context) error {
		query := c.QueryParams()
		fmt.Println("query", query)
		return c.String(200, "success")
	})

	userApi := api.Group("/user")

	userApi.POST("/register", controllers.RegisterUser)
	userApi.POST("/login", controllers.Login)
	supplierApi := api.Group("/suppliers")
	supplierApi.Use(middleware.Authentication())
	supplierApi.GET("/sync-products", controllers.SyncProductSupplier)
	supplierApi.GET("/products", controllers.GetProductSupplier)

	e.Logger.Fatal(e.Start(":8080"))
}
