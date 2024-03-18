package usecase

import "ngx/port"

type userUsecase struct {
	property usecaseProperty
}

func NewUserUsecase(property usecaseProperty) port.UserUsecase {
	return &userUsecase{property: property}
}
