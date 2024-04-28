package controller

import (
	"BookAPI/model"
	"BookAPI/usecase"
	"net/http"
	"strconv"

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
	bu usecase.IBookUsecase
}

func NewBookController(bu usecase.IBookUsecase) IBookController {
	return &bookController{bu}	
}

func (bc *bookController) GetAllBooks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token) // デコード
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	booksRes, err := bc.bu.GetAllBooks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, booksRes)
}

func (bc *bookController) GetBookById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("bookId")
	bookId, _  := strconv.Atoi(id)
	bookRes, err := bc.bu.GetBookById(uint(userId.(float64)), uint(bookId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bookRes)
}

func (bc *bookController) CreateBook(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	book := model.Book{}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	book.UserId = uint(userId.(float64))
	bookRes, err := bc.bu.CreateBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, bookRes)
}

func (bc *bookController) UpdateBook(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("bookId")
	bookId, _ := strconv.Atoi(id)

	book := model.Book{}
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	bookRes, err := bc.bu.UpdateBook(book, uint(userId.(float64)), uint(bookId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bookRes)
}