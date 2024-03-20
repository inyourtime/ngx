package usecase

import (
	"ngx/port"
	"ngx/util"
	"ngx/util/token"
)

type usecaseProperty struct {
	config     util.Config
	tokenMaker token.Maker
	repo       port.Repository
	logger     port.Logger
}

type usecases struct {
	property usecaseProperty
	authUc   port.AuthUsecase
	userUc   port.UserUsecase
}

func New(config util.Config, repo port.Repository, logger port.Logger) (port.Usecase, error) {
	tokenMaker, err := token.NewJwtMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	property := usecaseProperty{
		config:     config,
		repo:       repo,
		tokenMaker: tokenMaker,
		logger:     logger,
	}
	uc := usecases{
		property: property,
		authUc:   NewAuthUsecase(property),
		userUc:   NewUserUsecase(property),
	}
	return &uc, nil
}

func (u *usecases) Auth() port.AuthUsecase {
	return u.authUc
}

func (u *usecases) User() port.UserUsecase {
	return u.userUc
}

func (u *usecases) TokenMaker() token.Maker {
	return u.property.tokenMaker
}
