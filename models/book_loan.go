package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type UsersBook struct {
	gorm.Model
	UsersID     uint         `json:"users_id"`
	Users       Users        `gorm:"foreignKey:UsersID;reference:ID"`
	BookID      uint         `json:"book_id"`
	Book        Book         `gorm:"foreignKey:BookID;reference:ID"`
	Description string       `json:"description"`
	ReturnAt    sql.NullTime `json:"returnAt"`
}

type ResponseGetBook_loan struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	UsersName   string    `json:"usersName"`
	BookTitle   string    `json:"bookTitle"`
	CreatedAt   time.Time `json:"createdAt"`
	ReturnAt    time.Time `json:"returnAt"`
}

type ResponseUsersBook struct {
	UsersID     uint         `json:"users_id"`
	BookID      uint         `json:"book_id"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"createdAt"`
	ReturnAt    sql.NullTime `json:"returnAt"`
}

type RequestBook_loan struct {
	UsersID     uint   `json:"users_id"`
	BookID      uint   `json:"book_id"`
	Description string `json:"description"`
}

type ResponseBook_loan struct {
	ID          uint         `json:"id"`
	UsersID     uint         `json:"users_id"`
	BookID      uint         `json:"book_id"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"createdAt"`
	ReturnAt    sql.NullTime `json:"returnAt"`
}
