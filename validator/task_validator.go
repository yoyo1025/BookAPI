package validator

import "BookAPI/model"

type IBookValidator interface {
	BookValidate(book model.Book)
}

type BookValidator struct{}