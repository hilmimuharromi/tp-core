package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"tp-core/server/dto"
	"tp-core/server/integrations"
	"tp-core/server/repositories"
)

func SyncProductSupplier(c echo.Context) error {
	spCode := c.Param("id")
	if spCode == "" {
		return echo.NewHTTPError(400, "invalid id supplier")
	}
	log.Println("masuuuukkk 17")
	supplier, errSuplier := repositories.GetSupplier(spCode)
	log.Println("suppliers ===>", supplier)
	if errSuplier != nil {
		log.Println("errors suppliers ===>", errSuplier.Error())
		return echo.NewHTTPError(500, errSuplier.Error())
	}
	dataProducts, errFetch := integrations.FetchproductDs(supplier)
	if errFetch != nil {
		log.Println("errors fetch products ===>", errFetch.Error())
		return echo.NewHTTPError(500, errFetch.Error())
	}
	resSave, err := repositories.InsertManyProductSuppliers(dataProducts)
	if err != nil {
		log.Println("error", err)
		return err
	}
	res := map[string]interface{}{
		"data": resSave,
	}
	return c.JSON(200, res)
}

func GetProductSupplier(c echo.Context) error {
	//user := c.Get("user").(*jwt.Token)
	//claims := user.Claims.(*helpers.JwtCustomClaims)
	var param dto.ParamGetSupplierProduct
	if err := c.Bind(&param); err != nil {
		fmt.Println("error login===>", err)
		return echo.NewHTTPError(400, err.Error())
	}
	fmt.Println("param ===> 37", param)

	resData, err := repositories.GetSupplierProduct(param)
	if err != nil {
		log.Println("error", err)
	}

	//log.Println("resss ====>", resData)

	return c.JSON(200, resData)
}
