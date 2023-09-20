package models

import (
	"go-movies-crud/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publicaiton"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.CreateBook(&b)
	return b
}
func GetBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books

}
