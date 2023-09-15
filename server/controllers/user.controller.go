package controllers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"tp-core/server/dto"
	"tp-core/server/helpers"
	"tp-core/server/models"
	"tp-core/server/repositories"
)

func RegisterUser(c echo.Context) error {
	var user dto.RegisterDTO
	if err := c.Bind(&user); err != nil {
		fmt.Println("error ===>", err)
		return echo.NewHTTPError(400, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	passHash, _ := helpers.GenerateHash(user.Password)
	data := models.User{
		Email:    user.Email,
		Password: passHash,
		Id:       uuid.NewString(),
		Name:     user.Name,
	}

	resData, err := repositories.Register(data)

	if err != nil {
		return err
	}

	token, _ := helpers.CreateJwtToken(resData.Id)

	res := map[string]interface{}{
		"data":  resData,
		"token": token,
	}
	return c.JSON(200, res)
}

func Login(c echo.Context) error {
	var user dto.LoginDTO
	if err := c.Bind(&user); err != nil {
		fmt.Println("error login===>", err)
		return echo.NewHTTPError(400, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	userData := models.User{
		Email: user.Email,
	}

	resRepo, err := repositories.Login(userData)

	if err != nil {
		return err
	}

	if resRepo.Id == "" {
		return echo.NewHTTPError(400, "Please Check your email/password")
	}

	token, _ := helpers.CreateJwtToken(resRepo.Id)

	resRepo.Token = token

	res := map[string]interface{}{
		"data": resRepo,
	}
	return c.JSON(200, res)

}
