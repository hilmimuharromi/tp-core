package config

import (
	"fmt"
	"log"
	"os"
	models2 "tp-core/server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DB() *gorm.DB {
	return database
}

//var (
//	db *gorm.DB = DatabaseInit()
//)

func DatabaseInit() *gorm.DB {
	//host := "localhost"
	//user := "root"
	//password := "1234"
	//dbName := "go_rest_apis d"
	//port := 5432
	LoadEnv()

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		log.Println(e)
		panic(e)
	}
	return database
}

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&models2.Supplier{},
		&models2.ProductSupplier{},
		&models2.User{},
	)
	if err != nil {
		fmt.Println("============ Database Migration is Error ============")
	} else {
		fmt.Println("============ Database Migration Completed ============")
	}
}
