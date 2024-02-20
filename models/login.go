package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     uint32 `json:"role" validate:"required"`
	User     User
}

type RequestLogin struct {
	LoginUsername   string `json:"login_username" validate:"required"`
	LoginEmail      string `json:"login_email" validate:"required"`
	LoginPassword   string `json:"login_password" validate:"required"`
	LoginRole       uint32 `json:"login_role"`
	UserName        string `json:"user_name"`
	UserPhoneNumber string `json:"user_phonenumber"`
	UserAddress     string `json:"user_address"`
}

type GetLoginResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Addresss    string `json:"address"`
	Email       string `json:"email"`
	Role        uint32 `json:"role"`
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (login *Login) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	login.Password = string(bytes)

	return login.Password, nil
}

func (login *Login) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
