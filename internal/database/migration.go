package database

import (
	"log"

	"github.com/laurati/dolar_exchange_rate/internal/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&entity.ExchangeRate{},
	); err != nil {
		log.Fatalln("Error when automigrate db: " + err.Error())
	}
	log.Println("Database migration completed!")
}
