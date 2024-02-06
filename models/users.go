package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name         string      `json:"name"`
	Address      string      `json:"address"`
	Email        string      `json:"email"`
	Phone_number string      `json:"phone_number"`
	RoleID       uint        `json:"role_id"`
	Role         Role        `json:"role"`
	Books        []UsersBook `json:"books"`
}

type UsersResponse struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
	RoleID       uint   `json:"role_id"`
}

type GetUsersResponse struct {
	ID           uint         `json:"id"`
	Name         string       `json:"name"`
	Address      string       `json:"address"`
	Email        string       `json:"email"`
	Phone_number string       `json:"phone_number"`
	RoleID       uint         `json:"role_id"`
	Role         RoleResponse `json:"role"`
}
