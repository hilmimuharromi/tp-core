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
	totalOfElements  int `json:"totalOfElements"`
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

func FetchproductDs() []models.ProductSupplier {
	url := "http://36.88.177.95:8080/api/produk/pricing/0?page=0&size=10000&sort=namaoperator,ASC&sort=hargajual,asc"
	header := map[string]string{
		"Auth": "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJPSzAwNDAiLCJpc3MiOiJJUlNYIiwiaWF0IjoxNjk0Njg3MzIxLCJleHAiOjE2OTQ3MzA1MjB9.oGG6FlJjVeZaDPfusP5JUMzhRAX_DDT8J-3D3Dnzms8",
	}
	response := helpers.ClientRequest(http.MethodGet, url, "", header)

	var f ResProductDS
	json.Unmarshal(response, &f)

	myMap := f
	contents := myMap.Data.Content

	var dataProducts []models.ProductSupplier
	for idx, val := range contents {
		p := models.ProductSupplier{
			ID:         fmt.Sprintf("%s-%s-%d", "ds", val.Kodeproduk, idx),
			Name:       val.Namaproduk,
			Price:      val.Hargajual,
			SupplierId: 1,
			Operator:   val.Namaoperator,
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

	log.Println("responseee", myMap)
	return dataProducts
}
