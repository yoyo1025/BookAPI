package validator

import "BookAPI/model"

type IBookValidator interface {
	BookValidate(book model.Book)
}