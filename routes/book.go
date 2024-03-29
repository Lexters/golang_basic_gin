package routes

import (
	"golang_project_2024/config"
	"golang_project_2024/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetBook(c *gin.Context) {
	books := []models.Book{}

	config.DB.Preload("Author").Find(&books)

	GetBookResponse := []models.GetBookResponse{}

	for _, b := range books {
		author := models.AuthorResponse{
			ID:   b.AuthorID,
			Name: b.Author.Name,
		}

		bks := models.GetBookResponse{
			BookID:    b.ID,
			Title:     b.Title,
			Publisher: b.Publisher,
			Years:     b.Years,
			Author:    author,
		}
		GetBookResponse = append(GetBookResponse, bks)
	}

	c.JSON(200, gin.H{
		"Message": "Welcome Book!",
		"data":    GetBookResponse,
	})
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	data := config.DB.Preload(clause.Associations).First(&book, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	bk := models.GetBookResponse{
		BookID:    book.ID,
		Title:     book.Title,
		Publisher: book.Publisher,
		Years:     book.Years,
		Author: models.AuthorResponse{
			Name: book.Author.Name,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    bk,
	})

}

func PostBook(c *gin.Context) {
	// validate := validator.New()
	reqBook := models.Book{}
	c.BindJSON(&reqBook)

	// errs := validate.Struct(&reqBook)
	// if errs != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Bad request",
	// 	})

	// 	c.Abort()
	// 	return
	// }

	config.DB.Create(&reqBook)

	c.JSON(200, gin.H{
		"Message": "Insert Successfully!",
		"data":    reqBook,
	})
}

func PutBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	var reqBook models.Book
	c.BindJSON(&reqBook)

	config.DB.Model(&book).Where("id = ?", id).Updates(reqBook)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully!",
		"data":    reqBook,
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	books := models.Book{}

	config.DB.Delete(&books, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete Successfully",
	})
}
