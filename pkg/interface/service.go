package _interface

import (
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/service/request"
	"github.com/purwalenta/purwalenta/pkg/service/response"
)

type UserService interface {
	Login(echo.Context, request.UserLogin) (response.UserLogin, error)
	SignUp(echo.Context, request.UserSignUp) (response.UserSignUp, error)
}

type StudentService interface {

}

type MentorService interface {
}
