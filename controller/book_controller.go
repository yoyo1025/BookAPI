package controller

import (
	"BookAPI/model"
	"BookAPI/usecase"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IBookController interface {
	GetAllBooks(c echo.Context) error
	GetBookById(c echo.Context) error
	CreateBook(c echo.Context) error
	UploadPicture(c echo.Context) error
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	books, err := bc.bu.GetAllBooks(uint(userId))
	if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for i, book := range books {
			pictures, err := bc.bu.GetPicturesByBookId(book.ID)
			if err != nil {
					return c.JSON(http.StatusInternalServerError, err.Error())
			}
			for j, pic := range pictures {
					encodedImage := base64.StdEncoding.EncodeToString(pic.Image)
					pictures[j].ImageBase64 = encodedImage // Assume ImageBase64 field exists
			}
			books[i].Pictures = pictures
	}

	return c.JSON(http.StatusOK, books)
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

func (bc *bookController) UploadPicture(c echo.Context) error {
	// ユーザー情報の取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// リクエストからファイルを取得
	form, err := c.MultipartForm()
	if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to read multipart form data: "+err.Error())
	}
	files := form.File["image"]
	if len(files) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "No image file found in the request")
	}
	file, err := files[0].Open()
	if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open image file: "+err.Error())
	}
	defer file.Close()

	// ファイルの内容を読み込む
	imgData, err := ioutil.ReadAll(file)
	if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read image file: "+err.Error())
	}

	picture := model.Picture{
		Image:    imgData,
		UserId:   uint(userId.(float64)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
}

pictureRes, err := bc.bu.UploadPicture(picture)
if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
}

encodedImage := base64.StdEncoding.EncodeToString(picture.Image)
return c.JSON(http.StatusOK, model.PictureResponse{
		ID:          pictureRes.ID,
		ImageBase64: encodedImage,
})
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

func (bc *bookController) DeleteBook(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("bookId")
	bookId, _ := strconv.Atoi(id)

	err := bc.bu.DeleteBook(uint(userId.(float64)), uint(bookId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}