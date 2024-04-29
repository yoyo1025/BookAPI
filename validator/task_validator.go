package validator

import "BookAPI/model"

type IBookValidator interface {
	BookValidate(book model.Book)
}

type bookValidator struct{}

func NewBookValidator() IBookValidator {
	return &bookValidator{}
}