package database

import (
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/entity"
)

type UserRepository struct {
	DB interface{}
}

func (o *UserRepository) Login(ctx echo.Context, user entity.User) (*entity.User, error) {
	return nil, nil
}

func (o *UserRepository) SignUp(ctx echo.Context, user entity.User) (*entity.User, error) {
	return nil, nil
}
