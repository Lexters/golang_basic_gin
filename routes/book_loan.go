package routes

import (
	"database/sql"
	"golang_project_2024/config"
	"golang_project_2024/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetBookLoan(c *gin.Context) {
	UsersBook := []models.UserBook{}

	config.DB.Preload(clause.Associations).Find(&UsersBook)

	responseGetBook_loan := []models.ResponseGetBook_loan{}

	for _, ub := range UsersBook {
		rgb := models.ResponseGetBook_loan{
			Id:          ub.ID,
			Description: ub.Description,
			UserName:    ub.User.Name,
			BookTitle:   ub.Book.Title,
			CreatedAt:   ub.CreatedAt,
			ReturnAt:    ub.ReturnAt.Time,
		}

		responseGetBook_loan = append(responseGetBook_loan, rgb)
	}

	c.JSON(200, gin.H{
		"Message": "Welcome Book_loan!",
		"data":    responseGetBook_loan,
	})
}

func PostBookLoanByUsersID(c *gin.Context) {
	var reqBook_loan models.RequestBook_loan

	if err := c.ShouldBindJSON(&reqBook_loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	book_loan := models.UserBook{
		UserID:      reqBook_loan.UserID,
		BookID:      reqBook_loan.BookID,
		Description: reqBook_loan.Description,
	}

	insert := config.DB.Create(&book_loan)
	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   insert.Error.Error(),
		})

		c.Abort()
		return
	}

	respBook_loan := models.ResponseBook_loan{
		ID:          book_loan.BookID,
		UserID:      book_loan.UserID,
		BookID:      book_loan.BookID,
		Description: book_loan.Description,
		CreatedAt:   book_loan.CreatedAt,
	}

	c.JSON(200, gin.H{
		"Message": "Book_loan by Users ID!",
		"data":    respBook_loan,
	})
}

func GetBookLoanByBookID(c *gin.Context) {
	id := c.Param("id")

	Books := models.Book{}
	usBook := []models.ResponseUserBook{}

	data := config.DB.Preload(clause.Associations).Find(&Books, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	for _, book := range Books.User {

		usBook = append(usBook, models.ResponseUserBook{
			UserID:      book.UserID,
			BookID:      book.BookID,
			Description: book.Description,
			CreatedAt:   book.CreatedAt,
		})
	}

	respBook := models.ResponseBookUser{
		BookTitle: Books.Title,
		BookYears: Books.Years,
		UserBook:  usBook,
		Author: models.AuthorResponse{
			Name: Books.Author.Name,
		},
	}

	c.JSON(200, gin.H{
		"Message": "Welcome Book_loan by Book ID!",
		"data":    respBook,
	})
}

func UpdateBookLoanReturn(c *gin.Context) {
	id := c.Param("id")

	reqBLR := models.UserBook{}

	data := config.DB.First(&reqBLR, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	timeNow := time.Now()

	reqBLR.ReturnAt = sql.NullTime{Time: timeNow, Valid: true}

	config.DB.Model(&reqBLR).Where("id = ?", id).Updates(reqBLR)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully!",
		"data":    reqBLR,
	})

}

func DeleteBookLoan(c *gin.Context) {
	id := c.Param("id")

	bookloan := models.UserBook{}

	config.DB.Delete(&bookloan, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete Successfully",
	})
}
