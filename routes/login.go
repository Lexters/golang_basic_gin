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

	loginReqPass := models.Login{
		Password: reqRegis.LoginPassword,
	}

	//hash user password
	hash, err := loginReqPass.HashPassword(loginReqPass.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bad request",
			"errpr":   err.Error(),
		})

		c.Abort()
		return
	}

	//:: INITIALIZE DATA LOGIN
	Login := models.Login{
		Username: reqRegis.LoginUsername,
		Email:    reqRegis.LoginEmail,
		Password: hash,
		Role:     reqRegis.LoginRole,
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

	// insert login
	insertLogin := config.DB.Create(&Login)
	if insertLogin.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bad request",
			"error":   insertLogin.Error.Error(),
		})

		c.Abort()
		return
	}

	// Insert into user
	users := models.User{
		Name:         reqRegis.UserName,
		Address:      reqRegis.UserAddress,
		Phone_number: reqRegis.UserPhoneNumber,
		LoginID:      Login.ID,
	}

	if errs := validate.Struct(&users); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})

		c.Abort()
		return
	}

	if insertUser := config.DB.Create(&users); insertUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Insert User",
			"error":   insertLogin.Error.Error(),
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

func PutPassword(c *gin.Context) {
	// validate := validator.New()
	id := c.Param("id")

	data := config.DB.First(&models.Login{}, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	reqLog := models.RequestLogin{}
	c.BindJSON(&reqLog)

	loginReqPass := models.Login{
		Password: reqLog.LoginPassword,
	}

	//hash user password
	hash, err := loginReqPass.HashPassword(loginReqPass.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Bad request",
			"errpr":   err.Error(),
		})

		c.Abort()
		return
	}

	log := models.Login{
		Password: hash,
	}

	// if errs := validate.Struct(&log); errs != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Bad request",
	// 	})

	// 	c.Abort()
	// 	return
	// }

	config.DB.Model(&models.Login{}).Where("id = ?", id).Updates(&log)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully",
		"data":    log,
	})
}

func PutEmail(c *gin.Context) {
	// validate := validator.New()
	id := c.Param("id")

	data := config.DB.First(&models.Login{}, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	reqLog := models.RequestLogin{}
	c.BindJSON(&reqLog)

	log := models.Login{
		Email: reqLog.LoginEmail,
	}

	// if errs := validate.Struct(&reqLog); errs != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Bad request",
	// 	})

	// 	c.Abort()
	// 	return
	// }

	// check email
	checkEmail := config.DB.Where("email = ?", log.Email).First(&log)
	if checkEmail.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email Already Exists",
		})

		c.Abort()
		return
	}

	config.DB.Model(&models.Login{}).Where("id = ?", id).Updates(&log)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully",
		"data":    log,
	})
}

func PutUsername(c *gin.Context) {
	// validate := validator.New()
	id := c.Param("id")

	data := config.DB.First(&models.Login{}, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	reqLog := models.RequestLogin{}
	c.BindJSON(&reqLog)

	log := models.Login{
		Username: reqLog.LoginUsername,
	}

	// if errs := validate.Struct(&reqLog); errs != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Bad request",
	// 	})

	// 	c.Abort()
	// 	return
	// }

	config.DB.Model(&models.Login{}).Where("id = ?", id).Updates(&log)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Update Successfully",
		"data":    log,
	})
}

func DeleteLogin(c *gin.Context) {
	id := c.Param("id")

	login := models.Login{}
	users := models.User{}

	config.DB.Delete(&users, "id = ?", id)
	config.DB.Delete(&login, "id = ?", id)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Delete Successfully",
	})
}
