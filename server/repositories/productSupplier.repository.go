package repositories

import (
	"fmt"
	"strings"
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

func GetSupplierProduct(params dto.ParamGetSupplierProduct) (dto.ResGetSupplierDto, error) {
	limit := params.Limit
	page := params.Page
	category := strings.ToLower(params.Category)
	operator := strings.ToLower(params.Operator)
	name := strings.ToLower(params.Name)
	var resFormat dto.ResGetSupplierDto
	var data []dto.SupplierProductDto
	var count int64
	db := config.DB()
	dbQuery := db.Model(&models.ProductSupplier{})
	//.Count(&count).Limit(limit).Offset(offset)

	if limit != 0 {
		offset := page * limit
		dbQuery.Limit(limit).Offset(offset)
	}

	fmt.Println("masuuuuuuuuuuuuuuuuuuuuuuuuuuuk", params)

	if params.Category != "" {
		dbQuery.Where("category = ?", category)
	}

	if params.Operator != "" {
		dbQuery.Where("operator LIKE ?", "%"+operator+"%")
	}

	if params.Name != "" {
		dbQuery.Where("LOWER(name) LIKE ?", "%"+name+"%")
	}

	err := dbQuery.Count(&count).Find(&data).Error
	if err != nil {
		return resFormat, err
	}

	//fmt.Println("param ===>", limitInt, pageInt)
	//
	if limit == 0 {
		resFormat.Meta.Limit = int(count)
		resFormat.Meta.TotalPage = 1
	} else {
		resFormat.Meta.Limit = limit
		totalPage := int(count)/limit + 1
		resFormat.Meta.TotalPage = totalPage
	}
	//
	resFormat.Meta.Total = int(count)
	resFormat.Meta.Page = page
	resFormat.Meta.Limit = limit
	resFormat.Data = data
	resFormat.Code = 200
	return resFormat, err
}
