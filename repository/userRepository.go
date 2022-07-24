package repository

import (
	"registration/model"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface{
	FindByID(ID int) (model.User, error)
	Create(user model.User) (model.User, error)
	Login(email string) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(ID int) (model.User, error)  {
	var user model.User
	err := r.db.Find(&user, ID).Error
	return user, err
}

func (r *repository) Create(user model.User) (model.User, error)  {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) Login(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}