package validator

import (
	"BookAPI/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IBookValidator interface {
	BookValidate(book model.Book) error
}

type bookValidator struct{}

func NewBookValidator() IBookValidator {
	return &bookValidator{}
}

func (bv *bookValidator) BookValidate(book model.Book) error {
	return validation.ValidateStruct(&book,
		validation.Field(
			&book.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		))
}