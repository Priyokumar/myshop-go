package database

import (
	"log"
	"myshop/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB DB
var DB *gorm.DB

// InitiateDB InitiateDB
func InitiateDB() {

	dsn := "host=localhost user=test password=test dbname=testdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(

		&model.Item{},
		&model.QRCode{},
		&model.BarCode{},
		&model.Product{},
		&model.Purchase{},
		&model.PurchaseItem{},
	)

	if err != nil {
		log.Println(err.Error())
		return
	}

	DB = db

	log.Println("Initialized DB")

}
