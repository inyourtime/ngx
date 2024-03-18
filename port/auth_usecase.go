package port

import (
	"context"
	"ngx/domain"
)

type SignUpParams struct {
	User domain.User
}

type AuthUsecase interface {
	SignUp(context.Context, SignUpParams) (domain.User, error)
}
