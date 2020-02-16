package mail

import (
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/entity"
	"github.com/purwalenta/purwalenta/pkg/repository/mail/driver"
)

type UserMailingRepository struct {
	Driver driver.EmailDriver
}

func (mail *UserMailingRepository) SendSignUpVerification(ctx echo.Context, email entity.TemplateEmail) (bool, error) {
	err := driver.SendTemplateMail(mail.Driver, email)
	return nil == err, err
}

func (mail *UserMailingRepository) SendForgotPassword(ctx echo.Context, email entity.TemplateEmail) (bool, error) {
	err := driver.SendTemplateMail(mail.Driver, email)
	return nil == err, err
}
