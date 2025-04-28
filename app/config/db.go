package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDBConnection() *gorm.DB {
	once.Do(func() {
		var err error
		dsn := os.Getenv("DB_DSN")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Error connecting to database. Error:", err)
		}
	})

	return db
}
