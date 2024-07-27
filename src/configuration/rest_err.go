package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Função 'Error' que retorna uma string
// Isso é para satisfazer a interface de error do GoLang
func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message string, err string, code int, causes []Causes) *RestErr { //Esse *RestErr é um ponteiro que é o retorno da função
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnprocessableEntityError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unprocessable_entity",
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewUnprocessableEntityValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unprocessable_entity",
		Code:    http.StatusUnprocessableEntity,
		Causes:  causes,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}
