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

func (b *Book) CreateBook() (*Book, error){
	if db == nil {
		return nil, gorm.ErrInvalidDB
	}
	if err := db.Create(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllBooks() ([]Book, error){
	var books []Book
	if err := db.Find(&books).Error; err != nil{
		return nil, err
	}
	return books, nil
}

func GetBookById(id uint) (*Book, error) {
	var book Book
	if err := db.Find(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *Book) UpdateBook() (*Book, error){
	if err := db.Save(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func DeleteBook(id uint) error {
	if err := db.Delete(&Book{}, id).Error; err != nil {
		return err
	}
	return nil
}