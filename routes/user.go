package routes

import (
	"golang_project_2024/config"
	"golang_project_2024/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	Users := []models.Login{}
	config.DB.Preload("User").Find(&Users)

	getLoginResponse := []models.GetLoginResponse{}

	for _, u := range Users {
		post := models.GetLoginResponse{
			Name:        u.User.Name,
			PhoneNumber: u.User.Phone_number,
			Addresss:    u.User.Address,
			Email:       u.Email,
			Role:        u.Role,
		}

		getLoginResponse = append(getLoginResponse, post)
	}
	c.JSON(200, gin.H{
		"Message": "Welcome Users!",
		"data":    getLoginResponse,
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	users := models.Login{}

	data := config.DB.Preload("User").First(&users, "id = ?", id)

	//validate data
	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Users not found",
		})
		return
	}

	usr := models.GetLoginResponse{
		Name:        users.User.Name,
		PhoneNumber: users.User.Phone_number,
		Addresss:    users.User.Address,
		Email:       users.Email,
		Role:        users.Role,
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
		"data":    usr,
	})
}

func PutUser(c *gin.Context) {
	id := c.Param("id")

	var reqUsers models.User
	c.BindJSON(&reqUsers)

	config.DB.Model(&models.User{}).Where("id = ?", id).Updates(reqUsers)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully!",
		"data":    reqUsers,
	})
}
