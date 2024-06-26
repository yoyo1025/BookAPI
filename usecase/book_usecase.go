package usecase

import (
	"BookAPI/model"
	"BookAPI/repository"
	"BookAPI/validator"
)


type IBookUsecase interface {
	GetAllBooks(userId uint) ([]model.BookResponse, error)
	GetBookById(userId uint, bookId uint) (model.BookResponse, error)
	CreateBook(book model.Book) (model.BookResponse, error)
	UploadPicture(picture model.Picture) (model.PictureResponse, error)
	UpdateBook(book model.Book, userId uint, bookId uint) (model.BookResponse, error)
	DeleteBook(userId uint, bookId uint) error
	GetPicturesByBookId(bookId uint) ([]model.Picture, error)
}

type bookUsecase struct {
	br repository.IBookRepository
	bv validator.IBookValidator
}

func NewBookUsecase(br repository.IBookRepository, bv validator.IBookValidator) IBookUsecase {
	return &bookUsecase{br, bv}
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
			Comment: v.Comment,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resBooks = append(resBooks, t)
	}
	return resBooks, nil
}

func (bu *bookUsecase) GetBookById(userId uint, bookId uint) (model.BookResponse, error) {
	book := model.Book{}
	if err := bu.br.GetBookById(&book, userId, bookId); err != nil {
		return model.BookResponse{}, err
	}
	resBook := model.BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
	return resBook, nil
}

func (bu *bookUsecase) CreateBook(book model.Book) (model.BookResponse, error) {
	if err := bu.br.CreateBook(&book); err != nil {
		return model.BookResponse{}, err
	}
	resBook := model.BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
	return resBook, nil
}

func (bu *bookUsecase) UploadPicture(picture model.Picture) (model.PictureResponse, error) {
	if err := bu.br.UploadPicture(&picture); err != nil {
		return model.PictureResponse{}, err
	}
	resPicture := model.PictureResponse{
		ID:        picture.ID,
	}
	return resPicture, nil
}

func (bu *bookUsecase) UpdateBook(book model.Book, userId uint, bookId uint) (model.BookResponse, error) {
	if err := bu.br.UpdateBook(&book, userId, bookId); err != nil {
		return model.BookResponse{}, err
	}
	resBook := model.BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
	return resBook, nil
}

func (bu *bookUsecase) DeleteBook(userId uint, bookId uint) error {
	if err := bu.br.DeleteBook(userId, bookId); err != nil {
		return err
	}
	return nil
}

func (bu *bookUsecase) GetPicturesByBookId(bookId uint) ([]model.Picture, error) {
	return bu.br.GetPicturesByBookId(bookId)
}