// Code generated by MockGen. DO NOT EDIT.
// Source: .\repository\book_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	model "BookAPI/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIBookRepository is a mock of IBookRepository interface.
type MockIBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIBookRepositoryMockRecorder
}

// MockIBookRepositoryMockRecorder is the mock recorder for MockIBookRepository.
type MockIBookRepositoryMockRecorder struct {
	mock *MockIBookRepository
}

// NewMockIBookRepository creates a new mock instance.
func NewMockIBookRepository(ctrl *gomock.Controller) *MockIBookRepository {
	mock := &MockIBookRepository{ctrl: ctrl}
	mock.recorder = &MockIBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBookRepository) EXPECT() *MockIBookRepositoryMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockIBookRepository) CreateBook(book *model.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", book)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockIBookRepositoryMockRecorder) CreateBook(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockIBookRepository)(nil).CreateBook), book)
}

// DeleteBook mocks base method.
func (m *MockIBookRepository) DeleteBook(userId, bookId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", userId, bookId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockIBookRepositoryMockRecorder) DeleteBook(userId, bookId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockIBookRepository)(nil).DeleteBook), userId, bookId)
}

// GetAllBooks mocks base method.
func (m *MockIBookRepository) GetAllBooks(books *[]model.Book, userID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBooks", books, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAllBooks indicates an expected call of GetAllBooks.
func (mr *MockIBookRepositoryMockRecorder) GetAllBooks(books, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBooks", reflect.TypeOf((*MockIBookRepository)(nil).GetAllBooks), books, userID)
}

// GetBookById mocks base method.
func (m *MockIBookRepository) GetBookById(book *model.Book, userId, bookId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", book, userId, bookId)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetBookById indicates an expected call of GetBookById.
func (mr *MockIBookRepositoryMockRecorder) GetBookById(book, userId, bookId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockIBookRepository)(nil).GetBookById), book, userId, bookId)
}

// GetPicturesByBookId mocks base method.
func (m *MockIBookRepository) GetPicturesByBookId(bookId uint) ([]model.Picture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPicturesByBookId", bookId)
	ret0, _ := ret[0].([]model.Picture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPicturesByBookId indicates an expected call of GetPicturesByBookId.
func (mr *MockIBookRepositoryMockRecorder) GetPicturesByBookId(bookId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPicturesByBookId", reflect.TypeOf((*MockIBookRepository)(nil).GetPicturesByBookId), bookId)
}

// UpdateBook mocks base method.
func (m *MockIBookRepository) UpdateBook(book *model.Book, userId, bookId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", book, userId, bookId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockIBookRepositoryMockRecorder) UpdateBook(book, userId, bookId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockIBookRepository)(nil).UpdateBook), book, userId, bookId)
}

// UploadPicture mocks base method.
func (m *MockIBookRepository) UploadPicture(picture *model.Picture) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadPicture", picture)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadPicture indicates an expected call of UploadPicture.
func (mr *MockIBookRepositoryMockRecorder) UploadPicture(picture interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPicture", reflect.TypeOf((*MockIBookRepository)(nil).UploadPicture), picture)
}
