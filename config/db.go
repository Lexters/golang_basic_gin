package config

import (
	"golang_project_2024/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/golang_project?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	DB.AutoMigrate(&models.Book{}, &models.Users{}, &models.UsersBook{}, &models.Role{}, &models.Login{})
}
