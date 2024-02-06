package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string      `json:"title"`
	Author    string      `json:"author"`
	Publisher string      `json:"publisher"`
	Years     string      `json:"years"`
	Users     []UsersBook `json:"users"`
}

type GetBookResponse struct {
	BookID    uint   `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Years     string `json:"years"`
}

type ResponseBookUsers struct {
	BookTitle  string `json:"book_title"`
	BookAuthor string `json:"book_author"`
	BookYears  string `json:"book_years"`
	UsersBook  []ResponseUsersBook
}
