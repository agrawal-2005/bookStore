package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dsn := "prashant:GoyalPiyush123@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
