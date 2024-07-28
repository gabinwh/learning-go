package user_controller

import (
	"github.com/gabinwh/learning-go/src/configuration/validation"
	"github.com/gabinwh/learning-go/src/controller/model/request"
	"github.com/gabinwh/learning-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Create(context *gin.Context) {
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("There are some incorrect fields, error=%s", err.Error())
		errRest := validation.ValidateUserError(err)
		context.JSON(errRest.Code, errRest)
		return
	}

	userResponse := response.UserResponse{
		ID:    "1",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	context.JSON(http.StatusCreated, userResponse)
	return
}
