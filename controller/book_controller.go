package controller

import (
	"BookAPI/usecase"
	"net/http"

	"github.com/golang-jwt/jwt"
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

func (bc *bookController) GetAllBooks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token) // デコード
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	booksRes, err := bc.bc.GetAllBooks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, booksRes)
}