package models

import (
	"gorm.io/gorm"
	"github.com/agrawal-2005/bookStore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
}

func Migrate() {
    if db != nil {
        db.AutoMigrate(&Book{})
    }
}