package database

import (
	"log"

	"github.com/camping/config"
	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	config.LoadAppConfig()
}

func GetDB() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(config.AppConfig.ConnectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")

	return db, nil
}
