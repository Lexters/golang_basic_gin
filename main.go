package main

import (
	"golang_project_2024/config"
	"golang_project_2024/middlewares"
	"golang_project_2024/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/home", GetHome)

	route := r.Group("/")
	{
		login := route.Group("/login")
		{
			login.POST("/register", routes.RegisterUser)
			login.POST("/login", routes.GenerateToken)
		}

		books := route.Group("books").Use(middlewares.IsStaff())
		{
			books.GET("/", routes.GetBook)
			books.GET("/:id", routes.GetBookByID)
			books.POST("/", routes.PostBook)
			books.PUT("/:id", routes.PutBook)
			books.DELETE("/:id", routes.DeleteBook)
		}

		users := route.Group("users").Use(middlewares.IsAdmin())
		{
			users.GET("/", routes.GetUsers)
			users.GET("/:id", routes.GetUsersByID)
			users.POST("/", routes.PostUsers)
			users.PUT("/:id", routes.PutUsers)
			users.DELETE("/:id", routes.DeleteUsers)
		}

		role := route.Group("role").Use(middlewares.IsAdmin())
		{
			role.GET("/", routes.GetRole)
			role.GET("/:id", routes.GetRoleById)
			role.POST("/", routes.PostRole)
			role.PUT("/:id", routes.PutRole)
			role.DELETE("/:id", routes.DeleteRole)
		}

		book_loan := route.Group("book_loan").Use(middlewares.IsStaff())
		{
			book_loan.GET("/book/:id", routes.GetBookLoanByBookID)
			book_loan.POST("/", routes.PostBookLoanByUsersID)
			book_loan.PUT("/:id", routes.UpdateBookLoanReturn)
			book_loan.GET("/", routes.GetBookLoan)
		}
	}

	r.Run(":9090")
}

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Welcome Home!",
	})
}
