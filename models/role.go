package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Users []Users `json:"users"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type GetRoleResponse struct {
	ID    uint            `json:"id"`
	Name  string          `json:"name"`
	Code  string          `json:"code"`
	Users []UsersResponse `json:"users"`
}
