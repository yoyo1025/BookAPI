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

func NewBookUsecase(br repository.IBookRepository) IBookUsecase {
	return &bookUsecase{br}
}

func (bu *bookUsecase) GetAllBooks(userId uint) ([]model.BookResponse, error) {
	books := []model.Book{}
	if err := bu.br.GetAllBooks(&books, userId); err != nil {
		return nil, err
	}
	resBooks := []model.BookResponse{}
	for _, v := range books {
		t := model.BookResponse{
			ID: v.ID,
			Title: v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resBooks = append(resBooks, t)
	}
	return resBooks, nil
}