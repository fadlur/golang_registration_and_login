package request

import (
	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func (userRequest *UserRequest) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	userRequest.Password = string(bytes)
	return nil
}