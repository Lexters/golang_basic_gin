package main

import (
	"golang_project_2024/config"
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

		books := route.Group("books") //.Use(middlewares.IsStaff())
		{
			books.GET("/", routes.GetBook)
			books.GET("/:id", routes.GetBookByID)
			books.POST("/", routes.PostBook)
			books.PUT("/:id", routes.PutBook)
			books.DELETE("/:id", routes.DeleteBook)
		}

		author := route.Group("author")
		{
			author.GET("/", routes.GetAuthor)
			author.POST("/", routes.PostAuthor)
			author.PUT("/:id", routes.PutAuthor)
			author.DELETE("/:id", routes.DeleteAuthor)
		}

		users := route.Group("users") //.Use(middlewares.IsAdmin())
		{
			// 	users.GET("/", routes.GetUser)
			// 	users.GET("/:id", routes.GetUserByID)
			users.POST("/", routes.PostUser)
			users.PUT("/:id", routes.PutUser)
			users.DELETE("/:id", routes.DeleteUser)
		}

		book_loan := route.Group("book_loan") //.Use(middlewares.IsStaff())
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
