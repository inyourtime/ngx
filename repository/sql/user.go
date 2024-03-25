package sql

import (
	"context"
	"errors"
	"ngx/domain"
	"ngx/port"
	"ngx/util/exception"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	err := r.db.WithContext(ctx).Create(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.User{}, exception.New(exception.TypeValidation, "Email are already existing", err)
		}
		return domain.User{}, err
	}
	return u, nil
}
