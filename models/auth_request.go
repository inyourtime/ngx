package models

type AuthSignUpRequest struct {
	Email     string  `json:"email" validate:"required,email"`
	Password  *string `json:"password" validate:"required,min=8"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
