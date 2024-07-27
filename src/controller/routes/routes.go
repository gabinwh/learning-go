package routes

import (
	"github.com/gabinwh/learning-go/src/controller/user_controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/user/:id", user_controller.FindById)
	r.GET("/userByEmail/:email", user_controller.FindByEmail)
	r.POST("/user", user_controller.Create)
	r.PUT("/user/:id", user_controller.Update)
	r.DELETE("/user/:id", user_controller.Delete)
}
