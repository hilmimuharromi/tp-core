package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"tp-core/server/config"
	"tp-core/server/helpers"
)

func Authentication() echo.MiddlewareFunc {
	config.LoadEnv()
	secretKey := os.Getenv("SECRET_KEY")
	// initialize JWT middleware instance
	log.Println("secreet keey", secretKey)
	configJWT := echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helpers.JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
		Skipper: func(c echo.Context) bool {
			if strings.HasSuffix(c.Path(), "/login") {
				return true
			}
			return false
		},
		SuccessHandler: func(c echo.Context) {
			user, ok := c.Get("user").(*jwt.Token)
			if ok {
				claims := user.Claims.(*helpers.JwtCustomClaims)
				//custom loader https://echo.labstack.com/guide/context
				//appContext := c.(*common.AppContext)
				//user := &model.User{}
				//user.Name = claims.Name
				//appContext.User = user
				fmt.Println(" Claim ID : ")
				fmt.Println(claims.Id)
				fmt.Println("Standart Claim ID: ")
			}
		},
	})

	return configJWT
}
