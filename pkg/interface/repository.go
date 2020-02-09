package _interface

import (
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/entity"
)

type UserRepository interface {
	FindExistingUser(ctx echo.Context, user entity.User) (*entity.User, error)
	Login(ctx echo.Context, user entity.User) (*entity.User, error)
	SignUp(ctx echo.Context, user entity.User) (bool, error)
}

type UserMailingRepository interface {
	SendSignUpVerification(ctx echo.Context, email entity.TemplateEmail) (bool, error)
}

type StudentRepository interface {
}

type MentorRepository interface {
}
