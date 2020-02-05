package service

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
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
	return response.UserLogin{}, nil
}

func (service *UserService) SignUp(ctx echo.Context, req request.UserSignUp) (response.UserSignUp, error) {
	var resp = response.UserSignUp{}

	existingUser, err := service.Repo.FindUserForSignUp(ctx, entity.User{
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})

	if nil != err {
		return resp, err
	}

	if takenFields, isTaken := validation.ValidateUserSignUpTakenFields(existingUser); isTaken {
		resp.SignUpInfo.Success = false
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

	resp.User.ID = uuid.String()
	resp.User.FullName = req.FullName
	resp.User.Username = req.Username
	resp.User.Email = req.Email
	resp.User.PhoneNumber = req.PhoneNumber
	resp.User.Type = req.Type

	return resp, nil
}
