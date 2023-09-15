package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"tp-core/server/integrations"
	"tp-core/server/repositories"
)

func SyncProductSupplier(c echo.Context) error {
	dataProducts := integrations.FetchproductDs()
	resSave, err := repositories.InsertManyProductSuppliers(dataProducts)
	if err != nil {
		log.Println("error", err)
	}
	res := map[string]interface{}{
		"data": resSave,
	}
	return c.JSON(200, res)
}

func GetProductSupplier(c echo.Context) error {
	//user := c.Get("user").(*jwt.Token)
	//claims := user.Claims.(*helpers.JwtCustomClaims)
	param := c.QueryParams()
	var limit, page string
	if param["limit"] != nil {
		limit = param["limit"][0]
	}
	if param["page"] != nil {
		page = param["page"][0]
	}
	fmt.Println("param ===> 37", limit, page)

	resData, err := repositories.GetSupplierProduct(limit, page)
	if err != nil {
		log.Println("error", err)
	}

	//log.Println("resss ====>", resData)

	return c.JSON(200, resData)
}
