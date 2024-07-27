package user_controller

import (
	"fmt"
	"github.com/gabinwh/learning-go/src/configuration/validation"
	"github.com/gabinwh/learning-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"log"
)

func Create(context *gin.Context) {
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("There are some incorrect fields, error=%s", err.Error())
		errRest := validation.ValidateUserError(err)
		context.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
