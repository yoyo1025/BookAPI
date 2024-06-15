package usecase

import (
	"BookAPI/model"
	repository "BookAPI/repository/mock"
	"BookAPI/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBookUsecase_GetAllBooks(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := repository.NewMockIBookRepository(ctrl)
    uc := usecase.NewBookUsecase(mockRepo, nil)

    userId := uint(1)
    books := []model.Book{{Title: "Book 1"}, {Title: "Book 2"}}
    mockRepo.EXPECT().GetAllBooks(gomock.Any(), userId).SetArg(0, books).Return(nil)

    result, err := uc.GetAllBooks(userId)

    assert.NoError(t, err)
    assert.Len(t, result, 2)
    assert.Equal(t, "Book 1", result[0].Title)
}

func TestBookUsecase_GetBookById(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := repository.NewMockIBookRepository(ctrl)
    uc := usecase.NewBookUsecase(mockRepo, nil)

    bookId := uint(1)
    userId := uint(1)
    expectedBook := model.Book{ID: bookId, Title: "Existing Book"}
    mockRepo.EXPECT().GetBookById(gomock.Any(), userId, bookId).SetArg(0, expectedBook).Return(nil)

    result, err := uc.GetBookById(userId, bookId)

    assert.NoError(t, err)
    assert.Equal(t, "Existing Book", result.Title)
}

func TestBookUsecase_UpdateBook(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := repository.NewMockIBookRepository(ctrl)
    uc := usecase.NewBookUsecase(mockRepo, nil)

    book := model.Book{ID: 1, Title: "Updated Book"}
    userId, bookId := uint(1), uint(1)
    mockRepo.EXPECT().UpdateBook(&book, userId, bookId).Return(nil)

    result, err := uc.UpdateBook(book, userId, bookId)

    assert.NoError(t, err)
    assert.Equal(t, "Updated Book", result.Title)
}

func TestBookUsecase_DeleteBook(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := repository.NewMockIBookRepository(ctrl)
    uc := usecase.NewBookUsecase(mockRepo, nil)

    userId, bookId := uint(1), uint(1)
    mockRepo.EXPECT().DeleteBook(userId, bookId).Return(nil)

    err := uc.DeleteBook(userId, bookId)

    assert.NoError(t, err)
}

func TestBookUsecase_GetPicturesByBookId(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := repository.NewMockIBookRepository(ctrl)
    uc := usecase.NewBookUsecase(mockRepo, nil)

    pictures := []model.Picture{{ID: 1}}
    bookId := uint(1)
    mockRepo.EXPECT().GetPicturesByBookId(bookId).Return(pictures, nil)

    result, err := uc.GetPicturesByBookId(bookId)

    assert.NoError(t, err)
    assert.Len(t, result, 1)
    assert.Equal(t, uint(1), result[0].ID)
}

func TestBookUsecase_UploadPicture(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := repository.NewMockIBookRepository(ctrl)
    uc := usecase.NewBookUsecase(mockRepo, nil)

    picture := model.Picture{ID: 1}
    mockRepo.EXPECT().UploadPicture(&picture).Return(nil)

    result, err := uc.UploadPicture(picture)

    assert.NoError(t, err)
    assert.Equal(t, uint(1), result.ID)
}
