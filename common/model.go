package common

type User struct {
	Id       int32  `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"-" validate:"required,strength"`
	Role     string `json:"role" validate:"required,validateRole"`
	IsActive bool   `json:"is_active"`
}

type ErrorResponse struct {
	Message string            `json:"message"`
	Details map[string]string `json:"validationErrors"`
}
