package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `json:"book"`
}

type AuthorResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetAuthorResponse struct {
	ID    uint           `json:"id"`
	Name  string         `json:"name"`
	Books []BookResponse `json:"book"`
}
