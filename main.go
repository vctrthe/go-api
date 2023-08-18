package main

import (
	"go-api/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(192.168.100.150:3306)/api_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userInput := user.RegisterUserInput{}
	userInput.Name = "Test simpan 2"
	userInput.Email = "example@mail.com"
	userInput.Occupation = "test3"
	userInput.Password = "test"
	userService.RegisterUser(userInput)
}
