package usecase

import (
	"BookAPI/model"
	"BookAPI/repository"
	"BookAPI/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (LoginRes, error)
	GetUserInfo(id string) (model.UserResponse, error)
}

// usecaseはrepositoryインターフェースだけに依存する
type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

type LoginRes struct {
	TokenString string `json:"token_string"`
	Id int `json:"id"`
	Email string `json:"email"`
	UserName string `json:"user_name"`
}

// usecaseにrepositoryをDIするためのコンストラクタ
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		UserName: user.UserName, 
		Email: user.Email, 
		Password: string(hash),
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID: newUser.ID,
		Email: newUser.Email,	
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (LoginRes, error) {
	if err := uu.uv.UserValidate(user); err != nil {
			return LoginRes{}, err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
			return LoginRes{}, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
			return LoginRes{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": storedUser.ID,
			"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
			return LoginRes{}, err
	}
	return LoginRes{
			TokenString: tokenString,
			Id: int(storedUser.ID),
			Email: storedUser.Email,
			UserName: storedUser.UserName,
	}, nil
}

func (uu *userUsecase) GetUserInfo(id string) (model.UserResponse, error) {
	var user model.User
	if err := uu.ur.GetUserById(&user, id); err != nil {
			return model.UserResponse{}, err
	}

	resUser := model.UserResponse{
			ID: user.ID,
			Email: user.Email,
			UserName: user.UserName,
	}

	return resUser, nil
}
