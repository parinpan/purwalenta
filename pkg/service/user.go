package service

import (
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/parinpan/purwalenta/pkg/config"
	"github.com/parinpan/purwalenta/pkg/entity"
	"github.com/parinpan/purwalenta/pkg/errord"
	_interface "github.com/parinpan/purwalenta/pkg/interface"
	"github.com/parinpan/purwalenta/pkg/service/builder"
	"github.com/parinpan/purwalenta/pkg/service/request"
	"github.com/parinpan/purwalenta/pkg/service/response"
	"github.com/parinpan/purwalenta/pkg/service/validation"
	"github.com/parinpan/purwalenta/pkg/util"
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
		return resp, err
	}

	resp.LoginInfo.Success = true
	userLogin.Token, err = util.GenerateUserLoginToken(config.GetConfig(), *userLogin)

	if nil != err {
		err = errord.New(ctx, err)(errord.ErrGeneralOnCommonScenario, errord.Option{WriteLog: true})
		return resp, err
	}

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
		err = errord.New(ctx, err)(errord.ErrGeneralOnCommonScenario, errord.Option{WriteLog: true})
		return resp, err
	}

	if takenFields, isTaken := validation.ValidateUserSignUpTakenFields(req, *existingUser); isTaken {
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
		Username:    req.Email, // now we consider using email as user's username too
		Email:       req.Email,
		Password:    hashedPassword,
		PhoneNumber: req.PhoneNumber,
		Status:      entity.InactiveUser,
		Type:        req.Type,
	})

	if nil != err {
		err = errord.New(ctx, err)(errord.ErrGeneralOnCommonScenario, errord.Option{WriteLog: true})
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
		err = errord.New(ctx, err)(errord.ErrGeneralOnCommonScenario, errord.Option{WriteLog: true})
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
		err = errord.New(ctx, err)(errord.ErrGeneralOnCommonScenario, errord.Option{WriteLog: true})
		return resp, err
	}

	template := builder.UserSignUpVerificationEmailTemplate(verification)
	_, err = service.MailingRepo.SendSignUpVerification(ctx, template)

	if nil != err {
		err = errord.New(ctx, err)(errord.ErrGeneralOnCommonScenario, errord.Option{WriteLog: true})
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
		err = errord.New(ctx, err)(errord.ErrGeneralOnCommonScenario, errord.Option{WriteLog: true})
		return resp, err
	}

	if verification.Email != req.Email || verification.VerificationCode != req.Code {
		err = errord.New(ctx, nil)(errord.ErrInvalidCodeOnUserVerify, errord.Option{WriteLog: true})
		return resp, err
	}

	// modify user status to active status
	verified, err := service.Repo.Verify(ctx, verification.User)

	if nil != err || !verified {
		err = errord.New(ctx, err)(errord.ErrGeneralOnCommonScenario, errord.Option{WriteLog: true})
		return resp, err
	}

	resp.Success = true
	resp.User.LoginInfo.Success = true
	resp.User.ID = verification.User.ID
	resp.User.FullName = verification.User.FullName
	resp.User.Username = verification.User.Username
	resp.User.Email = verification.User.Email
	resp.User.PhoneNumber = verification.User.PhoneNumber
	resp.User.ProfilePicture = verification.ProfilePicture
	resp.User.Balance = verification.User.Balance
	resp.User.Token = verification.Token
	resp.User.Status = entity.ActiveUser
	resp.User.Type = verification.User.Type

	return resp, nil
}

func (service *UserService) ForgotPassword(ctx echo.Context, req request.UserForgotPassword) (response.UserForgotPassword, error) {
	var resp response.UserForgotPassword

	user, err := service.Repo.FindExistingUser(ctx, entity.User{Email: req.Email})
	if nil != err || user.ID == "" {
		err = errord.New(ctx, err)(errord.ErrNoMatchAccountOnUserForgotPassword, errord.Option{WriteLog: true})
		return resp, err
	}

	newPassword := util.GeneratePassword()
	hashedPassword, _ := util.HashPassword(newPassword)
	changed, err := service.Repo.ChangePassword(ctx, entity.User{ID: user.ID, Password: hashedPassword})

	if nil != err || !changed {
		return resp, err
	}

	template := builder.UserForgotPasswordTemplate(*user, newPassword)
	sent, err := service.MailingRepo.SendForgotPassword(ctx, template)

	if nil != err || !sent {
		err = errord.New(ctx, err)(errord.ErrGeneralOnCommonScenario, errord.Option{WriteLog: true})
		return resp, err
	}

	resp.Success = true
	resp.Message = entity.UserForgotPasswordMailSuccessfullySent

	return resp, nil
}

func (service *UserService) ChangePassword(ctx echo.Context, req request.UserChangePassword) (response.UserChangePassword, error) {
	var resp response.UserChangePassword

	user, err := service.Repo.FindExistingUser(ctx, entity.User{Email: req.Email})
	if nil != err {
		return resp, err
	}

	if !util.MatchPasswordHash(req.OldPassword, user.Password) {
		err = errord.New(ctx, err)(errord.ErrNoMatchPasswordOnUserChangePassword, errord.Option{WriteLog: true})
		return resp, err
	}

	newPassword, _ := util.HashPassword(req.NewPassword)
	changed, err := service.Repo.ChangePassword(ctx, entity.User{ID: user.ID, Password: newPassword})

	if nil != err || !changed {
		err = errord.New(ctx, err)(errord.ErrGeneralOnCommonScenario, errord.Option{WriteLog: true})
		return resp, err
	}

	resp.Success = true
	resp.Message = entity.UserPasswordChangedSuccessfully

	return resp, nil
}
