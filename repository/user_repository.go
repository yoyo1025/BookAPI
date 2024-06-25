package repository

import (
	"BookAPI/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	GetUserById(user *model.User, id string) error
	GetUserByName(user *model.User, username string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string ) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserById(user *model.User, id string) error {
	return ur.db.Where("id = ?", id).First(user).Error
}

func (ur *userRepository) GetUserByName(user *model.User, username string) error {
	return ur.db.Where("user_name = ?", username).First(user).Error
}