package models

type ProductCategories struct {
	Id   string `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type ProductTypes struct {
	Id   string `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
