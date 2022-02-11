package db

import (
	"auth-microservice/src/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var db *gorm.DB

func newDb(cfg *config.EnvConfig) *gorm.DB {
	log.Println("Start database Connection")
	username := cfg.GetDbUsername()
	password := cfg.GetDbPassword()
	dbName := cfg.GetDbName()
	dbHost := cfg.GetDbHost()
	dbPort := cfg.GetDbPort()

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", dbHost, username, dbName, dbPort, password)
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}
	return db
}

func GetDB(cfg *config.EnvConfig) *gorm.DB {
	if db == nil {
		db = newDb(cfg)
	}

	return db
}

func CloseDB(db *gorm.DB) {
	log.Println("Close database Connection")
	_ = db.Close()
}
