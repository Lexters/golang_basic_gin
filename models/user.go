package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string     `json:"name"`
	Address      string     `json:"address"`
	Phone_number string     `json:"phone_number"`
	Books        []UserBook `json:"books"`
	LoginID      uint       `json:"login_id"`
}

type UserResponse struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone_number string `json:"phone_number"`
}

type GetUserResponse struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone_number string `json:"phone_number"`
}
