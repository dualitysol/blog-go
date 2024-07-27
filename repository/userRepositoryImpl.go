package repository

import (
	"github.com/dualitysol/go-crud/dto/users/requests"
	"github.com/dualitysol/go-crud/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// CreateOne implements UserRepository.
func (u *UserRepositoryImpl) CreateOne(user models.User) (users []models.User, err error) {
	panic("unimplemented")
}

// DeleteById implements UserRepository.
func (u *UserRepositoryImpl) DeleteById(id uint) (success bool, err error) {
	panic("unimplemented")
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() (user []models.User, err error) {
	panic("unimplemented")
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(id uint) (user models.User, err error) {
	panic("unimplemented")
}

// FindOne implements UserRepository.
func (u *UserRepositoryImpl) FindOne(filter requests.FindUsers) (user models.User, err error) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(userPayload requests.UpdateUser) (user models.User, err error) {
	panic("unimplemented")
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}
