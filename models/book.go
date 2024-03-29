package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string     `json:"title" validate:"required"`
	AuthorID  uint       `json:"author_id" validate:"required"`
	Author    Author     `json:"author"`
	Publisher string     `json:"publisher" validate:"required"`
	Years     string     `json:"years" validate:"required"`
	User      []UserBook `json:"users"`
}

type GetBookResponse struct {
	BookID    uint           `json:"id"`
	Title     string         `json:"title"`
	Publisher string         `json:"publisher"`
	Years     string         `json:"years"`
	AuthorID  uint           `json:"author_id"`
	Author    AuthorResponse `json:"author"`
}

type BookResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Years     string `json:"years"`
}

type ResponseBookUser struct {
	BookTitle string         `json:"book_title"`
	BookYears string         `json:"book_years"`
	AuthorID  uint           `json:"author_id"`
	Author    AuthorResponse `json:"author"`
	UserBook  []ResponseUserBook
}
