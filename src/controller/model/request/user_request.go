package request

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Password string `json:"password" binding:"required,min=6,max=12,containsany=!?@#$%&*"`
	Email    string `json:"email" binding:"required,email"`
	Age      int8   `json:"age" binding:"required,min=1,max=140"` //O int8 já valida se é número
}
