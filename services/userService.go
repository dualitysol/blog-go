package services

import (
	"github.com/dualitysol/go-crud/dto/users/response"
)

type UserService interface {
	CreateAccount(username string, email string, password string) (userData response.UserData, token string, err error)
	Authenticate(username string, password string) (userData response.UserData, token string, err error)
	ForgotPassword(email string) (resultMsg string, err error)
	UpdatePassword(userId uint, password string) bool
	SaveAccountInfo(userId string, firstName string, lastName string, age int, gender string, address string, website string) bool
	GetAccount(userId uint) interface{}
}
