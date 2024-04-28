package controller

import (
	"BookAPI/usecase"

	"github.com/labstack/echo/v4"
)

type IBookController interface {
	GetAllBooks(c echo.Context) error
	GetBookById(c echo.Context) error
	CreateBook(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
}

type bookController struct {
	bc usecase.IBookUsecase
}

func NewBookController(bu usecase.IBookUsecase) IBookController {
	return &bookController{bu}	
}