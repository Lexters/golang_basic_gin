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

		author := route.Group("author").Use(middlewares.IsStaff())
		{
			author.GET("/", routes.GetAuthor)
			author.GET("/:id", routes.GetAuhtorByID)
			author.POST("/", routes.PostAuthor)
			author.PUT("/:id", routes.PutAuthor)
			author.DELETE("/:id", routes.DeleteAuthor)
		}

		user := route.Group("user").Use(middlewares.IsAdmin())
		{
			user.GET("/", routes.GetUser)
			user.GET("/:id", routes.GetUserByID)
			user.POST("/", routes.PostUser)
			user.PUT("/:id", routes.PutUser)
			user.DELETE("/:id", routes.DeleteUser)
		}

		book_loan := route.Group("book_loan").Use(middlewares.IsStaff())
		{
			book_loan.GET("/book/:id", routes.GetBookLoanByBookID)
			book_loan.GET("/", routes.GetBookLoan)
			book_loan.POST("/", routes.PostBookLoanByUsersID)
			book_loan.PUT("/:id", routes.UpdateBookLoanReturn)
			book_loan.DELETE("/:id", routes.DeleteBookLoan)
		}
	}

	r.Run(":9090")
}

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Welcome Home!",
	})
}
