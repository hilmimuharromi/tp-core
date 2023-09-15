package dto

type ResGetSupplierDto struct {
	Meta MetaPagination       `json:"meta"`
	Data []SupplierProductDto `json:"data"`
	Code int                  `json:"code"`
}

type MetaPagination struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	TotalPage int `json:"totalPage"`
	Limit     int `json:"limit"`
}

type SupplierProductDto struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Operator   string `json:"operator"`
	Type       string `json:"type"`
	Price      int    `json:"price"`
	Status     string `json:"status"`
	SupplierId int    `json:"supplierId"`
}
