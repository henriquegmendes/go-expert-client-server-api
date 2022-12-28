package database

import (
	"github.com/henriquegmendes/go-expert-client-server-api/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func ConnectToDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("can't connect to database. error: %s", err.Error())
	}

	err = db.AutoMigrate(&models.Exchange{})
	if err != nil {
		log.Fatalf("can't migrate structs to database. error: %s", err.Error())
	}

	return db
}
