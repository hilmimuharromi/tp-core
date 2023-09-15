package main

import (
	"gorm.io/gorm"
	config2 "tp-core/server/config"
	"tp-core/server/router"
)

var (
	db *gorm.DB = config2.DatabaseInit()
)

func main() {
	config2.DatabaseInit()
	config2.MigrateDatabase(db)
	router.InitRouter()
}
