package service

import (
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/config"
	"github.com/purwalenta/purwalenta/pkg/entity"
	"github.com/purwalenta/purwalenta/pkg/errord"
	_interface "github.com/purwalenta/purwalenta/pkg/interface"
	"github.com/purwalenta/purwalenta/pkg/service/builder"
	"github.com/purwalenta/purwalenta/pkg/service/request"
	"github.com/purwalenta/purwalenta/pkg/service/response"
	"github.com/purwalenta/purwalenta/pkg/service/validation"
	"github.com/purwalenta/purwalenta/pkg/util"
)

type UserService struct {
	Repo        _interface.UserRepository
	CacheRepo   _interface.UserCacheRepository
	MailingRepo _interface.UserMailingRepository
}

func (service *UserService) Login(ctx echo.Context, req request.UserLogin) (response.UserLogin, error) {
	var resp = response.UserLogin{}

	userLogin, err := service.Repo.Login(ctx, entity.User{
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})

	if nil != err {
		err = errord.New(ctx, err)(errord.ErrNoAccountMatchOnUserLogin, errord.Option{WriteLog: true})
		return resp, err
	}

	if !util.MatchPasswordHash(req.Password, userLogin.Password) {
		err = errord.New(ctx, err)(errord.ErrNoMatchPasswordOnUserLogin, errord.Option{WriteLog: true})
		return resp, nil
	}

	resp.LoginInfo.Success = true
	userLogin.Token, _ = util.GenerateUserLoginToken(config.GetConfig(), *userLogin)

	resp.ID = userLogin.ID
	resp.FullName = userLogin.FullName
	resp.Username = userLogin.Username
	resp.Email = userLogin.Email
	resp.PhoneNumber = userLogin.PhoneNumber
	resp.Balance = userLogin.Balance
	resp.Token = userLogin.Token
	resp.Type = userLogin.Type

	return resp, nil
}

func (service *UserService) SignUp(ctx echo.Context, req request.UserSignUp) (response.UserSignUp, error) {
	var resp = response.UserSignUp{}

	existingUser, err := service.Repo.FindExistingUser(ctx, entity.User{
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})

	if nil != err {
		err = errord.New(ctx, err)(errord.ErrFindExistingUserOnUserSignUp, errord.Option{WriteLog: true})
		return resp, err
	}

	if takenFields, isTaken := validation.ValidateUserSignUpTakenFields(*existingUser); isTaken {
		resp.SignUpInfo.TakenFields = takenFields
		resp.SignUpInfo.UserAlreadyExist = true
		takenFieldsString := strings.Join(takenFields, ", ")

		err = errord.New(ctx, nil)(
			errord.ErrFieldHasTakenOnUserSignUp,
			errord.Option{
				WriteLog:       true,
				FormatterValue: []interface{}{takenFieldsString},
			},
		)

		return resp, err
	}

	uuid, _ := uuid.NewUUID()
	hashedPassword, _ := util.HashPassword(req.Password)

	resp.SignUpInfo.Success, err = service.Repo.SignUp(ctx, entity.User{
		ID:          uuid.String(),
		FullName:    req.FullName,
		Username:    req.Username,
		Email:       req.Email,
		Password:    hashedPassword,
		PhoneNumber: req.PhoneNumber,
		Status:      entity.InactiveUser,
		Type:        req.Type,
	})

	if nil != err {
		err = errord.New(ctx, err)(errord.ErrUserCreationOnUserSignUp, errord.Option{WriteLog: true})
		return resp, err
	}

	resp.ID = uuid.String()
	resp.FullName = req.FullName
	resp.Username = req.Username
	resp.Email = req.Email
	resp.PhoneNumber = req.PhoneNumber
	resp.Status = entity.InactiveUser
	resp.Type = req.Type

	// send user sign up email verification
	go service.SendVerificationCode(ctx, request.UserSendVerificationCode{Email: req.Email})

	return resp, nil
}

func (service *UserService) SendVerificationCode(ctx echo.Context, req request.UserSendVerificationCode) (response.UserSendVerificationCode, error) {
	var resp response.UserSendVerificationCode

	user, err := service.Repo.FindExistingUser(ctx, entity.User{Email: req.Email})
	if nil != err {
		err = errord.New(ctx, err)(errord.ErrGeneralOnUserSendVerificationCode, errord.Option{WriteLog: true})
		return resp, err
	}

	accessToken, _ := util.GenerateUserLoginToken(config.GetConfig(), *user)
	verification := entity.SignUpVerification{
		User:             *user,
		Token:            accessToken,
		VerificationCode: util.GenerateNumberByDigitLen(4),
	}

	verification, err = service.CacheRepo.SetSignUpVerificationCode(ctx, verification)
	if nil != err {
		return resp, err
	}

	template := builder.UserSignUpVerificationEmailTemplate(verification)
	_, err = service.MailingRepo.SendSignUpVerification(ctx, template)

	if nil != err {
		return resp, err
	}

	resp.Success = true
	resp.Email = verification.User.Email
	resp.ExpiredAt = verification.ExpiredAt
	resp.Token = verification.Token

	return resp, nil
}

func (service *UserService) Verify(ctx echo.Context, req request.UserVerification) (response.UserVerification, error) {
	var resp response.UserVerification

	verification, err := service.CacheRepo.GetSignUpVerificationCode(ctx, entity.SignUpVerification{
		User: entity.User{Email: req.Email},
	})

	if nil != err {
		err = errord.New(ctx, err)(errord.ErrGeneralOnUserVerify, errord.Option{WriteLog: true})
		return resp, err
	}

	if verification.Email != req.Email || verification.VerificationCode != req.Code {
		err = errord.New(ctx, nil)(errord.ErrInvalidCodeOnUserVerify, errord.Option{WriteLog: true})
		return resp, err
	}

	// modify user status to active status
	verified, err := service.Repo.Verify(ctx, verification.User)

	if nil != err || !verified {
		err = errord.New(ctx, nil)(errord.ErrGeneralOnUserVerify, errord.Option{WriteLog: true})
		return resp, err
	}

	resp.Success = true
	resp.User.ID = verification.User.ID
	resp.User.FullName = verification.User.FullName
	resp.User.Username = verification.User.Username
	resp.User.Email = verification.User.Email
	resp.User.PhoneNumber = verification.User.PhoneNumber
	resp.User.ProfilePicture = verification.ProfilePicture
	resp.User.Balance = verification.Balance
	resp.User.Token = verification.User.Token
	resp.User.Status = entity.ActiveUser
	resp.User.Type = verification.User.Type

	return resp, nil
}
