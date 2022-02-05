package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"log"
)

var db *gorm.DB

func Init() *gorm.DB {
	log.Println("Start database Connection")
	username := viper.Get("db.username")
	password := viper.Get("db.password")
	dbName := viper.Get("db.dbname")
	dbHost := viper.Get("db.host")
	dbPort := viper.Get("db.port")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", dbHost, username, dbName, dbPort, password)
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}
	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		db = Init()
	}

	return db
}

func CloseDB(db *gorm.DB){
	log.Println("Close database Connection")
	_ = db.Close()
}