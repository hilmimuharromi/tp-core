package repositories

import (
	"tp-core/server/config"
	"tp-core/server/dto"
	"tp-core/server/models"
)

func Register(user models.User) (models.User, error) {
	db := config.DB()
	err := db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

func Login(user models.User) (dto.ResLogin, error) {
	var res dto.ResLogin
	db := config.DB()
	err := db.Model(&models.User{}).Where("email = ?", user.Email).Find(&user).Error
	if err != nil {
		return res, err
	}

	res.Id = user.Id
	res.Name = user.Name
	res.Email = user.Email
	return res, err
}
