package repository

import (
	"BookAPI/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBookRepository interface {
	GetAllBooks(books *[]model.Book, userID uint) error
	GetBookById(book *model.Book, userId uint, bookId uint) error
	CreateBook(book *model.Book) error
	UploadPicture(picture *model.Picture) error
	UpdateBook(book *model.Book, userId uint, bookId uint) error
	DeleteBook(userId uint, bookId uint) error
	GetPicturesByBookId(bookId uint) ([]model.Picture, error)
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

func (br *bookRepository) CreateBook(book *model.Book) error {
	if err := br.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) UploadPicture(picture *model.Picture) error {
	if err := br.db.Create(picture).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookRepository) UpdateBook(book *model.Book, userId uint, bookId uint) error {
	result := br.db.Model(book).Clauses(clause.Returning{}).Where("id=? AND user_id=?", bookId, userId).Update("title", book.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exsit")
	}
	return nil
}

func (br *bookRepository) DeleteBook(userId uint, bookId uint) error {
	result := br.db.Where("id=? AND user_id=?", bookId, userId).Delete(&model.Book{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (br *bookRepository) GetPicturesByBookId(bookId uint) ([]model.Picture, error) {
	var pictures []model.Picture
	if err := br.db.Where("id = ?", bookId).Find(&pictures).Error; err != nil {
		return nil, err
	}
	return pictures, nil
}