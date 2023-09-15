package models

import "gorm.io/gorm"

type ProductSupplier struct {
	gorm.Model
	ID         string `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Operator   string `json:"operator"`
	Type       string `json:"type"`
	Price      int    `json:"price"`
	Status     string `json:"status"`
	SupplierId int    `json:"supplierId"`
	Supplier   Supplier
}
