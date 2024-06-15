package usecase

import (
	"BookAPI/model"
	repository "BookAPI/repository/mock"
	"BookAPI/usecase"
	validator "BookAPI/validator/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockIUserRepository(ctrl)
	mockValidator := validator.NewMockIUserValidator(ctrl)

	u := usecase.NewUserUsecase(mockRepo, mockValidator)

	// テストユーザーの設定
	testUser := model.User{
			Email:    "test@example.com",
			Password: "password123",
	}

	// バリデーションが成功することを期待
	mockValidator.EXPECT().UserValidate(gomock.Any()).Return(nil)

	// パスワードハッシュとユーザーの作成が呼ばれることを期待
	mockRepo.EXPECT().CreateUser(gomock.Any()).DoAndReturn(func(user *model.User) error {
			assert.NotEqual(t, "", user.Password, "Password should be hashed")
			user.ID = 1 // 新しいIDを設定
			return nil
	})

	// 実際にSignUpを呼び出し
	response, err := u.SignUp(testUser)

	// 結果の検証
	assert.NoError(t, err)
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "test@example.com", response.Email)
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockIUserRepository(ctrl)
	mockValidator := validator.NewMockIUserValidator(ctrl)

	u := usecase.NewUserUsecase(mockRepo, mockValidator)

	testUser := model.User{
			Email:    "test@example.com",
			Password: "password123",
	}

	// バリデーションが成功することを期待
	mockValidator.EXPECT().UserValidate(gomock.Any()).Return(nil)

	// GetUserByEmailが正しいユーザーを返すことを期待
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	mockRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).SetArg(0, model.User{
			ID:       1,
			Email:    "test@example.com",
			Password: string(hashedPassword),
	})

	// 実際にLoginを呼び出し
	token, err := u.Login(testUser)

	// 結果の検証
	assert.NoError(t, err)
	assert.NotEmpty(t, token, "Token should not be empty")
}
