package integrations

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"tp-core/server/helpers"
	"tp-core/server/models"
)

type ResProductDS struct {
	Success    bool          `json:"success"`
	StatusCode string        `json:"statusCode"`
	Msg        string        `json:"msg"`
	Data       DataproductDs `json:"data"`
}

type DataproductDs struct {
	Number           int `json:"number"`
	NumberOfElements int `json:"numberOfElements"`
	TotalOfElements  int `json:"totalOfElements"`
	Page             int `json:"page"`
	TotalPage        int `json:"totalPage"`
	Content          []ContentDs
}

type ContentDs struct {
	Id           int    `json:"id"`
	Kodeproduk   string `json:"kodeproduk"`
	Namaproduk   string `json:"namaproduk"`
	Keterangan   string `json:"keterangan"`
	Namaoperator string `json:"namaoperator"`
	Operatorimg  string `json:"operatorimg"`
	ProdukImg    string `json:"produkImg"`
	Nominal      int    `json:"nominal"`
	Hargajual    int    `json:"hargajual"`
	Poin         int    `json:"poin"`
	Jenisproduk  string `json:"jenisproduk"`
	Aktif        string `json:"aktif"`
	Isgangguan   int    `json:"isgangguan"`
	Isstokkosong int    `json:"isstokkosong"`
}

func FetchproductDs(suplier models.Supplier) []models.ProductSupplier {
	url := suplier.PriceUrl
	header := map[string]string{
		"Auth": fmt.Sprintf("Bearer %s", suplier.Token),
	}
	response := helpers.ClientRequest(http.MethodGet, url, "", header)

	var f ResProductDS
	err := json.Unmarshal(response, &f)
	if err != nil {
		return nil
	}

	myMap := f
	contents := myMap.Data.Content

	var dataProducts []models.ProductSupplier
	for idx, val := range contents {
		p := models.ProductSupplier{
			ID:         fmt.Sprintf("%s-%s-%d", "ds", val.Kodeproduk, idx),
			Name:       val.Namaproduk,
			Price:      val.Hargajual,
			SupplierId: 1,
			Operator:   helpers.GetOperator(val.Namaoperator),
			Category:   helpers.GetCategory(val.Namaproduk, val.Namaoperator),
			//Operator: val.Namaoperator,
		}
		//if strings.Contains(strings.ToLower(val.Namaproduk), "data") ||
		//	strings.Contains(strings.ToLower(val.Namaoperator), "inject") ||
		//	strings.Contains(strings.ToLower(val.Namaoperator), "by.u") {
		//	p.Category = "DATA"
		//} else if strings.Contains(strings.ToLower(val.Namaproduk), "reguler") || strings.Contains(strings.ToLower(val.Namaoperator), "reguler") {
		//	p.Category = "PULSA"
		//} else if strings.Contains(strings.ToLower(val.Namaproduk), "transfer") || strings.Contains(strings.ToLower(val.Namaoperator), "transfer") {
		//	p.Category = "PULSA TRANSFER"
		//}

		if val.Isgangguan == 1 {
			p.Status = "gangguan"
		} else if val.Isstokkosong == 1 {
			p.Status = "kosong"
		} else {
			p.Status = "aktif"
		}
		dataProducts = append(dataProducts, p)
	}

	log.Println("responseee", myMap)
	return dataProducts
}
