package service

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/purwalenta/purwalenta/pkg/config"
	"github.com/purwalenta/purwalenta/pkg/entity"
	_interface "github.com/purwalenta/purwalenta/pkg/interface"
	"github.com/purwalenta/purwalenta/pkg/service/request"
	"github.com/purwalenta/purwalenta/pkg/service/response"
	"github.com/purwalenta/purwalenta/pkg/service/validation"
	"github.com/purwalenta/purwalenta/pkg/util"
)

type UserService struct {
	Repo _interface.UserRepository
}

func (service *UserService) Login(ctx echo.Context, req request.UserLogin) (response.UserLogin, error) {
	var resp = response.UserLogin{}

	userLogin, err := service.Repo.Login(ctx, entity.User{
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})

	if nil != err {
		return resp, err
	}

	if !util.MatchPasswordHash(req.Password, userLogin.Password) {
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
		return resp, err
	}

	if takenFields, isTaken := validation.ValidateUserSignUpTakenFields(*existingUser); isTaken {
		resp.SignUpInfo.TakenFields = takenFields
		resp.SignUpInfo.UserAlreadyExist = true
		return resp, nil
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
		Type:        req.Type,
	})

	if nil != err {
		return resp, err
	}

	resp.ID = uuid.String()
	resp.FullName = req.FullName
	resp.Username = req.Username
	resp.Email = req.Email
	resp.PhoneNumber = req.PhoneNumber
	resp.Type = req.Type

	return resp, nil
}
