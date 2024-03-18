package models

type AuthSignUpRequest struct {
	Email    string  `json:"email" validate:"required,email"`
	Name     *string `json:"name"`
	Password *string `json:"password" validate:"required,min=8"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
