package services

import (
	"github.com/dualitysol/go-crud/dto/users/response"
	"github.com/dualitysol/go-crud/models"
	"github.com/dualitysol/go-crud/repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	validate       *validator.Validate
}

func (u *UserServiceImpl) CreateAccount(username string, email string, password string) (userData response.UserData, token string, err error) {
	payload := models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	r := u.UserRepository.CreateOne(payload)

	helper.PanicError(r.Error)

	return r
}

func (u *UserServiceImpl) Authenticate(username string, password string) (userData response.UserData, token string, err error) {

}

func (u *UserServiceImpl) ForgotPassword(email string) (resultMsg string, err error) {

}

func (u *UserServiceImpl) UpdatePassword(userId uint, password string) bool {

}

func (u *UserServiceImpl) SaveAccountInfo(userId string, firstName string, lastName string, age int, gender string, address string, website string) bool {

}

func (u *UserServiceImpl) GetAccount(userId uint) interface{} {

}
