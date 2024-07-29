package user_controller

import (
	rest_err "github.com/gabinwh/learning-go/src/configuration"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) Find(context *gin.Context) {
	err := rest_err.NewBadRequestError("Rota errada!")

	context.JSON(err.Code, err)
}

func (uc *userControllerInterface) FindByEmail(context *gin.Context) {

}
