package main

import (
	"gojwt/app/http/controllers"
	"gojwt/app/models"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	api := r.Group("/api")
	models.ConnectDB()
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)
	r.Run(":8180")

}
