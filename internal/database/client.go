package database

import (
	"log"

	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(
		&models.Fund{},
		&models.Bot{},
		&models.Exchange{},
		&models.TradingBalance{},
		&models.ExchangeBalance{},
		&models.ExchangeCurrencyBalance{},
		&models.Contributions{},
	)
	log.Println("Database Migration Completed!")
}
