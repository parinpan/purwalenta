package _interface

import (
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/entity"
)

type UserRepository interface {
	FindExistingUser(ctx echo.Context, user entity.User) (*entity.User, error)
	Login(ctx echo.Context, user entity.User) (*entity.User, error)
	SignUp(ctx echo.Context, user entity.User) (bool, error)
	Verify(ctx echo.Context, user entity.User) (bool, error)
}

type UserMailingRepository interface {
	SendSignUpVerification(ctx echo.Context, email entity.TemplateEmail) (bool, error)
}

type UserCacheRepository interface {
	GetSignUpVerificationCode(ctx echo.Context, verification entity.SignUpVerification) (entity.SignUpVerification, error)
	SetSignUpVerificationCode(ctx echo.Context, verification entity.SignUpVerification) (entity.SignUpVerification, error)
}

type OauthRepository interface {
	GetUserInfo(ctx echo.Context, oauth entity.Oauth) (*entity.User, error)
}

type StudentRepository interface {
}

type MentorRepository interface {
}
