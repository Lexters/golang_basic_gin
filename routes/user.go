package routes

import (
	"golang_project_2024/config"
	"golang_project_2024/models"
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

// func GetUserByID(c *gin.Context) {
// 	id := c.Param("id")

// 	users := models.Login{}

// 	data := config.DB.Preload("User").First(&users, "id = ?", id)

// 	//validate data
// 	if data.Error != nil {
// 		log.Printf(data.Error.Error())
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "Users not found",
// 		})
// 		return
// 	}

// 	usr := models.GetLoginResponse{
// 		Name:        users.User.Name,
// 		PhoneNumber: users.User.Phone_number,
// 		Addresss:    users.User.Address,
// 		Email:       users.Email,
// 		Role:        users.Role,
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"Message": "Success",
// 		"data":    usr,
// 	})
// }

// func GetUser(c *gin.Context) {
// 	Users := []models.User{}
// 	config.DB.Preload("Role").Find(&Users)

// 	getUserResponse := []models.GetUserResponse{}

// 	for _, u := range Users {
// 		Role := models.RoleResponse{
// 			ID:   u.Role.ID,
// 			Name: u.Role.Name,
// 			Code: u.Role.Code,
// 		}

// 		usrs := models.GetUserResponse{
// 			ID:           u.ID,
// 			Name:         u.Name,
// 			Address:      u.Address,
// 			Email:        u.Email,
// 			Phone_number: u.Phone_number,
// 		}

// 		getUserResponse = append(getUserResponse, usrs)
// 	}

// 	c.JSON(200, gin.H{
// 		"Message": "Welcome Users!",
// 		"data":    getUserResponse,
// 	})
// }

// func GetUserByID(c *gin.Context) {
// 	id := c.Param("id")

// 	users := models.User{}

// 	data := config.DB.Preload("Role").First(&users, "id = ?", id)

// 	if data.Error != nil {
// 		log.Printf(data.Error.Error())
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "Users not found",
// 		})
// 		return
// 	}

// 	role := models.RoleResponse{
// 		ID:   users.Role.ID,
// 		Name: users.Role.Name,
// 		Code: users.Role.Code,
// 	}

// 	getUsersResponse := models.GetUserResponse{
// 		ID:           users.ID,
// 		Name:         users.Name,
// 		Address:      users.Address,
// 		Email:        users.Email,
// 		Phone_number: users.Phone_number,
// 		RoleID:       users.RoleID,
// 		Role:         role,
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"Message": "Success",
// 		"data":    getUsersResponse,
// 	})

// }

func PostUser(c *gin.Context) {
	reqUsers := models.User{}
	c.BindJSON(&reqUsers)

	config.DB.Create(&reqUsers)

	c.JSON(200, gin.H{
		"Message": "Insert Successfully!",
		"data":    reqUsers,
	})
}

func PutUser(c *gin.Context) {
	id := c.Param("id")
	var users models.User

	var reqUsers models.User
	c.BindJSON(&reqUsers)

	config.DB.Model(&users).Where("id = ?", id).Updates(reqUsers)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully!",
		"data":    reqUsers,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	users := models.User{}

	config.DB.Delete(&users, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete Successfully",
	})
}
