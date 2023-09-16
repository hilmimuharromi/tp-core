package web

import (
	"embed"
	"log"
	"net/url"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	//go:embed all:build
	build embed.FS
	//go:embed build/index.html
	indexHTML embed.FS

	distDirFS     = echo.MustSubFS(build, "build")
	distIndexHTML = echo.MustSubFS(indexHTML, "build")
)

func RegisterHandlers(e *echo.Echo) {
	if os.Getenv("ENV") == "dev" {
		log.Println("Running in dev mode")
		setupDevProxy(e)
		return
	}
	log.Println("Running in prod mode")
	// Use the static assets from the dist directory
	e.FileFS("/", "index.html", distIndexHTML)
	e.StaticFS("/", distDirFS)
}

func setupDevProxy(e *echo.Echo) {
	urlWeb, err := url.Parse("http://localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	// Setep a proxy to the vite dev server on localhost:5173
	balancer := middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: urlWeb,
		},
	})
	e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: balancer,
		Skipper: func(c echo.Context) bool {
			// Skip the proxy if the prefix is /api
			return len(c.Path()) >= 4 && c.Path()[:4] == "/api"
		},
	}))
}
