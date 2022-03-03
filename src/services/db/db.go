package db

import (
	"auth-microservice/src/config"
	"auth-microservice/src/log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

const dbUri = "host=%s user=%s dbname=%s port=%s sslmode=disable password=%s"

func newDb(cfg *config.EnvConfig, logger *log.Logger) *gorm.DB {
	logger.InfoLog.Println("Start database Connection")
	url := fmt.Sprintf(dbUri,
		cfg.GetDbHost(),
		cfg.GetDbUsername(),
		cfg.GetDbName(),
		cfg.GetDbPort(),
		cfg.GetDbPassword(),
	)

	db, err := gorm.Open("postgres", url)
	if err != nil {
		logger.ErrorLog.Panic(err)
	}
	return db
}

func GetDB(cfg *config.EnvConfig, logger *log.Logger) *gorm.DB {
	if db == nil {
		db = newDb(cfg, logger)
	}
	logger.InfoLog.Println("Connecting to database...")

	return db
}

func CloseDB(db *gorm.DB, logger *log.Logger) {
	logger.InfoLog.Println("Close database Connection")
	if err := db.Close(); err != nil {
		logger.ErrorLog.Panic(err)
	}
}
