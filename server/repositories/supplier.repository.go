package repositories

import (
	"tp-core/server/config"
	"tp-core/server/models"
)

func GetSupplier(spCode string) (models.Supplier, error) {
	var supplier models.Supplier
	db := config.DB()
	err := db.Model(&models.Supplier{}).Where("sp_code = ?", spCode).Find(&supplier).Error
	if err != nil {
		return supplier, err
	}
	return supplier, err
}
