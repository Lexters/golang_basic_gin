package routes

import (
	"golang_project_2024/config"
	"golang_project_2024/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetRole(c *gin.Context) {
	roles := []models.Role{}

	config.DB.Preload("Users").Find(&roles)

	GetRoleResponse := []models.GetRoleResponse{}

	for _, r := range roles {

		users := []models.UsersResponse{}
		for _, u := range r.Users {
			usr := models.UsersResponse{
				Name:         u.Name,
				Address:      u.Address,
				Email:        u.Email,
				Phone_number: u.Phone_number,
				RoleID:       u.RoleID,
			}

			users = append(users, usr)
		}

		rk := models.GetRoleResponse{
			ID:    r.ID,
			Name:  r.Name,
			Code:  r.Code,
			Users: users,
		}

		GetRoleResponse = append(GetRoleResponse, rk)
	}

	c.JSON(200, gin.H{
		"Message": "Welcome Roles!",
		"data":    GetRoleResponse,
	})
}

func GetRoleById(c *gin.Context) {
	id := c.Param("id")

	var role models.Role

	data := config.DB.Preload(clause.Associations).First(&role, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	rl := models.GetRoleResponse{
		ID:   role.ID,
		Name: role.Name,
		Code: role.Code,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    rl,
	})

}

func PostRole(c *gin.Context) {
	reqRole := models.Role{}
	c.BindJSON(&reqRole)

	config.DB.Create(&reqRole)

	c.JSON(200, gin.H{
		"Message": "Insert Successfully!",
		"data":    reqRole,
	})
}

func PutRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role

	var reqRole models.Role
	c.BindJSON(&reqRole)

	config.DB.Model(&role).Where("id = ?", id).Updates(reqRole)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully!",
		"data":    reqRole,
	})
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")

	roles := models.Role{}

	config.DB.Delete(&roles, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete Successfully",
	})
}
