package routes

import (
	"golang_project_2024/config"
	"golang_project_2024/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm/clause"
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

func GetAuhtorByID(c *gin.Context) {
	id := c.Param("id")

	var author models.Author

	data := config.DB.Preload(clause.Associations).First(&author, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	books := []models.BookResponse{}
	for _, b := range author.Books {
		bk := models.BookResponse{
			ID:        b.ID,
			Title:     b.Title,
			Publisher: b.Publisher,
			Years:     b.Years,
		}

		books = append(books, bk)
	}

	getAuthorResponse := models.GetAuthorResponse{
		ID:    author.ID,
		Name:  author.Name,
		Books: books,
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
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
