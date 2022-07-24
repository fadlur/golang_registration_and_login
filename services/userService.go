package services

import (
	"registration/model"
	"registration/repository"
	"registration/request"
)

type UserServiceInterface interface {
	FindByID(ID int) (model.User, error)
	Create(userRequest request.UserRequest) (model.User, error)
	Login(email string) (model.User, error)
}

type userService struct {
	userRepositoryInterface repository.UserRepositoryInterface
}

func InitUserService(userRepositoryInterface repository.UserRepositoryInterface) *userService  {
	return &userService{userRepositoryInterface}
}

func (service *userService) FindByID(ID int) (model.User, error) {
	user, err := service.userRepositoryInterface.FindByID(ID)
	return user, err
}

func (service *userService) Create(userRequest request.UserRequest) (model.User, error) {
	// var userRequest request.UserRequest
	var userReturn model.User
	if err := userRequest.HashPassword(userRequest.Password); err != nil {
		return userReturn, err
	}

	user := model.User{
		FirstName: userRequest.FirstName,
		LastName: userRequest.LastName,
		Email: userRequest.Email,
		Password: userRequest.Password,
	}

	newUser, err := service.userRepositoryInterface.Create(user)
	return newUser, err
}

func (service *userService) Login(email string) (model.User, error) {
	user, err := service.userRepositoryInterface.Login(email)
	return user, err
}