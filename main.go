package main

import (
	"tp-core/router"
)

func main() {
	// Create HTTP server

	//// Connect To Database
	//config.DatabaseInit()
	//gorm := config.DB()
	//
	//dbGorm, err := gorm.DB()
	//if err != nil {
	//	panic(err)
	//}
	//
	//dbGorm.Ping()

	router.InitRouter()

}
