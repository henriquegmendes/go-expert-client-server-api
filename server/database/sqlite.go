package database

import (
	"github.com/henriquegmendes/go-expert-client-server-api/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func ConnectToDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("can't connect to database. error: %s", err.Error())
	}

	err = db.AutoMigrate(&models.Exchange{})
	if err != nil {
		log.Fatalf("can't migrate structs to database. error: %s", err.Error())
	}

	return db
}
