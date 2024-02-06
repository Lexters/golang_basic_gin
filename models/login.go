package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     uint32 `json:"role" validate:"required"`
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (login *Login) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	login.Password = string(bytes)

	return nil
}

func (login *Login) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
