package repository

import (
	"BookAPI/model"

	"gorm.io/gorm"
)

type IBookRepository interface {
	GetAllBooks(books *[]model.Book, userID uint) error
	GetBookById(book *model.Book, userId uint, bookId uint) error
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book, userId uint, bookId uint) error
	DeleteBook(userId uint, bookId uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) GetAllBooks(books *[]model.Book, userId uint) error {
	if err := br.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(books).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) GetBookById(book *model.Book, userId uint, bookId uint) error {
	if err := br.db.Joins("User").Where("user_id=?", userId).First(book, bookId).Error; err != nil {
		return err 
	}
	return nil
}