package main

import (
	"github.com/gabinwh/learning-go/src/controller/routes"
	"github.com/gabinwh/learning-go/src/controller/user_controller"
	"github.com/gabinwh/learning-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserDomainService()
	userController := user_controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
