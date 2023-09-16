package repositories

import (
	"fmt"
	"strconv"
	"tp-core/server/config"
	"tp-core/server/dto"
	"tp-core/server/models"
)

//func AddSupplierProduct(productSupplier models.ProductSupplier) (models.ProductSupplier, error) {
//	db := config.DB()
//	err := db.Model(&models.ProductSupplier{}).Create(&productSupplier).Error
//	if err != nil {
//		return productSupplier, err
//	}
//	return productSupplier, err
//}

func InsertManyProductSuppliers(data []models.ProductSupplier) (interface{}, error) {
	db := config.DB()
	err := db.Model(&models.ProductSupplier{}).Save(&data).Error
	if err != nil {
		return data, err
	}
	return data, err
}

func GetSupplierProduct(limit string, page string) (dto.ResGetSupplierDto, error) {
	var resFormat dto.ResGetSupplierDto
	var data []dto.SupplierProductDto
	var count int64
	var limitInt int
	var pageInt int
	var offset int
	db := config.DB()

	fmt.Println("masuuuuuuuuuuuuuuuuuuuuuuuuuuuk", limit, page)

	if page != "" {
		pageInt, _ = strconv.Atoi(page)
	} else {
		pageInt = 0
	}

	if limit != "" && limit != "0" {
		limitInt, _ = strconv.Atoi(limit)
		offset = (pageInt - 1) * limitInt
	} else {
		limitInt = -1
	}

	err := db.Model(&models.ProductSupplier{}).Count(&count).Limit(limitInt).Offset(offset).Find(&data).Error
	if err != nil {
		return resFormat, err
	}

	fmt.Println("param ===>", limitInt, pageInt)

	if limitInt == -1 {
		resFormat.Meta.Limit = int(count)
		resFormat.Meta.TotalPage = 1
	} else {
		resFormat.Meta.Limit = limitInt
		totalPage := int(count)/limitInt + 1
		resFormat.Meta.TotalPage = totalPage
	}

	resFormat.Meta.Total = int(count)
	resFormat.Meta.Page = pageInt
	resFormat.Data = data
	resFormat.Code = 200
	return resFormat, err
}
