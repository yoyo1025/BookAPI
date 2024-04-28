package repository

import "BookAPI/model"

type IBookRepository interface {
	GetAllBooks(books *[]model.Book, userID uint) error
	GetBookById(book *model.Book, userId uint, bookId uint) error
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book, userId uint, bookId uint) error
	DeleteBook(userId uint, bookId uint) error
}