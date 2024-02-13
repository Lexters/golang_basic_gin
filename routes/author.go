package routes

import (
	"golang_project_2024/config"
	"golang_project_2024/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func GetAuthor(c *gin.Context) {
	author := []models.Author{}

	config.DB.Preload("Books").Find(&author)

	getAuthorResponse := []models.GetAuthorResponse{}

	for _, a := range author {

		books := []models.BookResponse{}
		for _, b := range a.Books {
			bok := models.BookResponse{
				ID:        b.ID,
				Title:     b.Title,
				Publisher: b.Publisher,
				Years:     b.Years,
			}

			books = append(books, bok)
		}

		auth := models.GetAuthorResponse{
			ID:    a.ID,
			Name:  a.Name,
			Books: books,
		}

		getAuthorResponse = append(getAuthorResponse, auth)
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Welcome Author",
		"data":    getAuthorResponse,
	})
}

func PostAuthor(c *gin.Context) {
	validate := validator.New()
	reqAut := models.Author{}
	c.BindJSON(&reqAut)

	errs := validate.Struct(&reqAut)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})

		c.Abort()
		return
	}

	config.DB.Create(&reqAut)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Insert Successfully",
		"data":    reqAut,
	})
}

func PutAuthor(c *gin.Context) {
	id := c.Param("id")
	var auth models.Author

	var reqAuth models.Author
	c.BindJSON(&reqAuth)

	config.DB.Model(&auth).Where("id = ?", id).Updates(reqAuth)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully!",
		"data":    reqAuth,
	})
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")

	auth := models.Author{}

	config.DB.Delete(&auth, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete Successfully",
	})
}
