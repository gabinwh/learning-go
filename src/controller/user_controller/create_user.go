package user_controller

import (
	"github.com/gabinwh/learning-go/src/configuration/logger"
	"github.com/gabinwh/learning-go/src/configuration/validation"
	"github.com/gabinwh/learning-go/src/controller/model/request"
	"github.com/gabinwh/learning-go/src/model"
	"github.com/gabinwh/learning-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//var (
//	UserDomainInterface model.UserDomainInterface
//)

func (uc *userControllerInterface) Create(context *gin.Context) {
	logger.Info("Init Create User", zap.String("func", "CreateUser"))

	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao validar os campos.", err)
		errRest := validation.ValidateUserError(err)
		context.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Password,
		userRequest.Email,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		logger.Error("Erro ao criar user domain.", err)
		context.JSON(err.Code, err)
		return
	}

	logger.Info("Success Create User", zap.String("func", "CreateUser"))

	context.JSON(http.StatusCreated, view.ConvertDomainToResponse(domain))
}
