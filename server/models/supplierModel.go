package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name       string `json:"name"`
	SpCode     string `json:"spCode"`
	OurBalance int    `json:"ourBalance"`
	Token      string `json:"token"`
	PriceUrl   string `json:"priceUrl"`
}
