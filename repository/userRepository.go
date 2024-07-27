package repository

import (
	"github.com/dualitysol/go-crud/dto/users/requests"
	"github.com/dualitysol/go-crud/models"
)

type UserRepository interface {
	CreateOne(user models.User) (users []models.User, err error)
	FindAll() (user []models.User, err error)
	FindById(id uint) (user models.User, err error)
	FindOne(filter requests.FindUsers) (user models.User, err error)
	Update(userPayload requests.UpdateUser) (user models.User, err error)
	DeleteById(id uint) (success bool, err error)
}
