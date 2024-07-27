package user_controller

import (
	rest_err "github.com/gabinwh/learning-go/src/configuration"
	"github.com/gin-gonic/gin"
)

func FindById(c *gin.Context) {
	err := rest_err.NewBadRequestError("Rota errada!")

	c.JSON(err.Code, err)
}

func FindByEmail(context *gin.Context) {

}
