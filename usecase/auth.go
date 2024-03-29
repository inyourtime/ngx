package usecase

import (
	"context"
	"ngx/domain"
	"ngx/port"
)

type authUsecase struct {
	property usecaseProperty
}

func NewAuthUsecase(property usecaseProperty) port.AuthUsecase {
	return &authUsecase{property: property}
}

func (u *authUsecase) SignUp(ctx context.Context, arg port.SignUpParams) (domain.User, error) {
	user, err := domain.NewUser(arg.User)
	if err != nil {
		return domain.User{}, err
	}

	newUser, err := u.property.repo.User().Create(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}
