package validator

import "BookAPI/model"

type IUserValidator interface {
	UserValidte(user model.User) error
}