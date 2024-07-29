package routes

import (
	"github.com/gabinwh/learning-go/src/controller/user_controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController user_controller.UserControllerInterface) {

	r.GET("/user/:id", userController.Find)
	r.GET("/userByEmail/:email", userController.FindByEmail)
	r.POST("/user", userController.Create)
	r.PUT("/user/:id", userController.Update)
	r.DELETE("/user/:id", userController.Delete)
}
