package validation

import (
	"encoding/json"
	"errors"
	rest_err "github.com/gabinwh/learning-go/src/configuration"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	br_translation "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	Validate  = validator.New()
	translate ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {

		br := pt_BR.New()
		uni := ut.New(br)
		translate, _ = uni.GetTranslator("pt_BR")
		br_translation.RegisterDefaultTranslations(val, translate) //Lidar com o erro dessa função
	}
}

func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewUnprocessableEntityError("Campo(s) inválido(s).")
	} else if errors.As(validation_err, &jsonValidationError) {
		var errorCauses []rest_err.Causes

		for _, err := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: err.Translate(translate),
				Field:   err.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewUnprocessableEntityValidationError("Campo(s) inválido(s)", errorCauses)
	} else {
		return rest_err.NewInternalServerError("Internal Server Error")
	}

}
