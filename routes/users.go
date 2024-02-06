package routes

import (
	"golang_project_2024/config"
	"golang_project_2024/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	Users := []models.Users{}
	config.DB.Preload("Role").Find(&Users)

	getUserResponse := []models.GetUsersResponse{}

	for _, u := range Users {
		Role := models.RoleResponse{
			ID:   u.Role.ID,
			Name: u.Role.Name,
			Code: u.Role.Code,
		}

		usrs := models.GetUsersResponse{
			ID:           u.ID,
			Name:         u.Name,
			Address:      u.Address,
			Email:        u.Email,
			Phone_number: u.Phone_number,
			RoleID:       u.RoleID,
			Role:         Role,
		}

		getUserResponse = append(getUserResponse, usrs)
	}

	c.JSON(200, gin.H{
		"Message": "Welcome Users!",
		"data":    getUserResponse,
	})
}

func GetUsersByID(c *gin.Context) {
	id := c.Param("id")

	users := models.Users{}

	data := config.DB.Preload("Role").First(&users, "id = ?", id)

	if data.Error != nil {
		log.Printf(data.Error.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Users not found",
		})
		return
	}

	role := models.RoleResponse{
		ID:   users.Role.ID,
		Name: users.Role.Name,
		Code: users.Role.Code,
	}

	getUsersResponse := models.GetUsersResponse{
		ID:           users.ID,
		Name:         users.Name,
		Address:      users.Address,
		Email:        users.Email,
		Phone_number: users.Phone_number,
		RoleID:       users.RoleID,
		Role:         role,
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
		"data":    getUsersResponse,
	})

}

func PostUsers(c *gin.Context) {
	reqUsers := models.Users{}
	c.BindJSON(&reqUsers)

	config.DB.Create(&reqUsers)

	c.JSON(200, gin.H{
		"Message": "Insert Successfully!",
		"data":    reqUsers,
	})
}

func PutUsers(c *gin.Context) {
	id := c.Param("id")
	var users models.Users

	var reqUsers models.Users
	c.BindJSON(&reqUsers)

	config.DB.Model(&users).Where("id = ?", id).Updates(reqUsers)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully!",
		"data":    reqUsers,
	})
}

func DeleteUsers(c *gin.Context) {
	id := c.Param("id")

	users := models.Users{}

	config.DB.Delete(&users, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete Successfully",
	})
}
