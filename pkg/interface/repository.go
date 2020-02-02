package _interface

import (
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/entity"
)

type UserRepository interface {
	Login(ctx echo.Context, user entity.User) (*entity.User, error)
	SignUp(ctx echo.Context, user entity.User) (*entity.User, error)
}

type StudentRepository interface {
}

type MentorRepository interface {
}
