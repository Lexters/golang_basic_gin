package routes

import (
	"golang_project_2024/auth"
	"golang_project_2024/config"
	"golang_project_2024/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RegisterUser(c *gin.Context) {
	validate := validator.New()
	var user models.Login

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	if errs := validate.Struct(&user); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})

		c.Abort()
		return
	}

	// check email
	checkEmail := config.DB.Where("email = ?", user.Email).First(&user)
	if checkEmail.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email Already Exists",
		})

		c.Abort()
		return
	}

	//hash user password
	err := user.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bad request",
			"errpr":   err.Error(),
		})

		c.Abort()
		return
	}

	// insert user
	insertUser := config.DB.Create(&user)
	if insertUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bad request",
			"error":   insertUser.Error.Error(),
		})

		c.Abort()
		return
	}

	//response register
	c.JSON(http.StatusCreated, gin.H{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}

func GenerateToken(c *gin.Context) {
	request := models.TokenRequest{}
	user := models.Login{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	// check email
	checkEmail := config.DB.Where("email = ?", request.Email).First(&user)
	if checkEmail.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Email Not Found",
			"error":   checkEmail.Error.Error(),
		})

		c.Abort()
		return
	}

	// check password
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password Not Match",
			"error":   credentialError.Error(),
		})

		c.Abort()
		return
	}

	// generate token
	tokenString, err := auth.GenerateJWT(user.Email, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
