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
	reqRegis := models.RequestLogin{}
	c.BindJSON(&reqRegis)

	Login := models.Login{
		Username: reqRegis.LoginUsername,
		Email:    reqRegis.LoginEmail,
		Password: reqRegis.LoginPassword,
		Role:     reqRegis.LoginRole,
		User: models.User{
			Name:         reqRegis.UserName,
			Address:      reqRegis.UserAddress,
			Phone_number: reqRegis.UserPhoneNumber,
		},
	}

	if err := c.ShouldBindJSON(&Login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})

		c.Abort()
		return
	}

	if errs := validate.Struct(&Login); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})

		c.Abort()
		return
	}

	// check email
	checkEmail := config.DB.Where("email = ?", Login.Email).First(&Login)
	if checkEmail.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email Already Exists",
		})

		c.Abort()
		return
	}

	//hash user password
	err := Login.HashPassword(Login.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bad request",
			"errpr":   err.Error(),
		})

		c.Abort()
		return
	}

	// insert user
	insertUser := config.DB.Create(&Login)
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
		"user_id":  Login.ID,
		"email":    Login.Email,
		"username": Login.Username,
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
