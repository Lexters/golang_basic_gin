package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type UserBook struct {
	gorm.Model
	UserID      uint         `json:"user_id"`
	User        User         `gorm:"foreignKey:UserID;reference:ID"`
	BookID      uint         `json:"book_id"`
	Book        Book         `gorm:"foreignKey:BookID;reference:ID"`
	Description string       `json:"description"`
	ReturnAt    sql.NullTime `json:"returnAt"`
}

type ResponseGetBook_loan struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	UserName    string    `json:"userName"`
	BookTitle   string    `json:"bookTitle"`
	CreatedAt   time.Time `json:"createdAt"`
	ReturnAt    time.Time `json:"returnAt"`
}

type ResponseUserBook struct {
	UserID      uint         `json:"user_id"`
	BookID      uint         `json:"book_id"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"createdAt"`
	ReturnAt    sql.NullTime `json:"returnAt"`
}

type RequestBook_loan struct {
	UserID      uint   `json:"user_id"`
	BookID      uint   `json:"book_id"`
	Description string `json:"description"`
}

type ResponseBook_loan struct {
	ID          uint         `json:"id"`
	UserID      uint         `json:"user_id"`
	BookID      uint         `json:"book_id"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"createdAt"`
	ReturnAt    sql.NullTime `json:"returnAt"`
}
