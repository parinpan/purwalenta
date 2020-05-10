package _interface

import (
	"github.com/labstack/echo"
	"github.com/parinpan/purwalenta/pkg/service/request"
	"github.com/parinpan/purwalenta/pkg/service/response"
)

type UserService interface {
	Login(echo.Context, request.UserLogin) (response.UserLogin, error)
	SignUp(echo.Context, request.UserSignUp) (response.UserSignUp, error)
	SendVerificationCode(echo.Context, request.UserSendVerificationCode) (response.UserSendVerificationCode, error)
	Verify(echo.Context, request.UserVerification) (response.UserVerification, error)
	ForgotPassword(echo.Context, request.UserForgotPassword) (response.UserForgotPassword, error)
	ChangePassword(echo.Context, request.UserChangePassword) (response.UserChangePassword, error)
}

type AssessmentService interface {
	FindPersonalityQuestions(echo.Context, request.FindPersonalityQuestions) (response.PersonalityQuestion, error)
}

type StudentService interface {
}

type MentorService interface {
}

type OauthService interface {
	Exchange(ctx echo.Context, req request.OauthExchange) (response.OauthExchange, error)
}
