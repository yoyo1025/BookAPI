package usecase

import (
	"BookAPI/model"
	"BookAPI/repository"
)


type IBookUsecase interface {
	GetAllBooks(userId uint) ([]model.BookResponse, error)
	GetBookById(userId uint, bookId uint) (model.BookResponse, error)
	CreateBook(book model.Book) (model.BookResponse, error)
	UpdateBook(book model.Book, userId uint, bookId uint) (model.BookResponse, error)
	DeleteBook(userId uint, bookId uint) error
}

type bookUsecase struct {
	br repository.IBookRepository
}