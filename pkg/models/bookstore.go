package models

import (
	"github.com/Omoalfa/bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func DBGetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) DBCreateBook() *Book {
	db.Create(&b)
	return b
}

func DBGetBookById(id int64) *Book {
	var book Book
	db.Where("ID=?", id).Find(&book)
	return &book
}

func DBDeleteBookById(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}

func DBUpdateBookById(id int64, data Book) Book {
	var book Book
	db.Model(&Book{}).Where("ID = ?", id).Updates(data)
	return book
}
