package integrations

import (
	"encoding/json"
	"errors"
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

func FetchproductDs(suplier models.Supplier) ([]models.ProductSupplier, error) {
	var dataProducts []models.ProductSupplier
	url := suplier.PriceUrl
	header := map[string]string{
		"Auth": fmt.Sprintf("Bearer %s", suplier.Token),
	}
	response := helpers.ClientRequest(http.MethodGet, url, "", header)
	log.Println("response get product ds", string(response))
	var f ResProductDS
	err := json.Unmarshal(response, &f)
	if err != nil {
		log.Error("error json unmarshal", string(response))
	}

	if f.Success != true {
		return []models.ProductSupplier{}, errors.New(fmt.Sprintf("error from suppliers : %s", f.Msg))
	}

	myMap := f
	contents := myMap.Data.Content

	for idx, val := range contents {
		p := models.ProductSupplier{
			ID:         fmt.Sprintf("%s-%s-%d", "ds", val.Kodeproduk, idx),
			Name:       val.Namaproduk,
			Price:      val.Hargajual,
			SupplierId: 1,
			Operator:   helpers.GetOperator(val.Namaoperator),
			Category:   helpers.GetCategory(val.Namaproduk, val.Namaoperator),
		}

		if val.Isgangguan == 1 {
			p.Status = "gangguan"
		} else if val.Isstokkosong == 1 {
			p.Status = "kosong"
		} else {
			p.Status = "aktif"
		}
		dataProducts = append(dataProducts, p)
	}

	return dataProducts, nil
}
