package handler

import (
	"github.com/labstack/echo"
	apiPkg "github.com/parinpan/purwalenta/pkg/api"
	"github.com/parinpan/purwalenta/pkg/delivery/rest/response"
	"github.com/parinpan/purwalenta/pkg/service/request"
)

type userHandler struct {
	api *apiPkg.UserAPI
}

func NewUserHandler() *userHandler {
	return &userHandler{
		api: apiPkg.NewUserAPI(apiPkg.DefaultUserAPIFlag),
	}
}

func (handler *userHandler) Login(ctx echo.Context) error {
	var req = new(request.UserLogin)
	var resp = response.New()
	var err error

	if err = ctx.Bind(req); nil != err {
		return resp.ToJSON(ctx)
	}

	if err = ctx.Validate(req); nil != err {
		return resp.ToJSON(ctx)
	}

	if resp.Data, err = handler.api.Service.Login(ctx, *req); nil != err {
		return resp.ToJSON(ctx)
	}

	return resp.ToJSON(ctx)
}

func (handler *userHandler) SignUp(ctx echo.Context) error {
	var req = new(request.UserSignUp)
	var resp = response.New()
	var err error

	if err = ctx.Bind(req); nil != err {
		return resp.ToJSON(ctx)
	}

	if err = ctx.Validate(req); nil != err {
		return resp.ToJSON(ctx)
	}

	if resp.Data, err = handler.api.Service.SignUp(ctx, *req); nil != err {
		return resp.ToJSON(ctx)
	}

	return resp.ToJSON(ctx)
}

func (handler *userHandler) Verify(ctx echo.Context) error {
	var req = new(request.UserVerification)
	var resp = response.New()
	var err error

	if err = ctx.Bind(req); nil != err {
		return resp.ToJSON(ctx)
	}

	if err = ctx.Validate(req); nil != err {
		return resp.ToJSON(ctx)
	}

	if resp.Data, err = handler.api.Service.Verify(ctx, *req); nil != err {
		return resp.ToJSON(ctx)
	}

	return resp.ToJSON(ctx)
}

func (handler *userHandler) ForgotPassword(ctx echo.Context) error {
	var req = new(request.UserForgotPassword)
	var resp = response.New()
	var err error

	if err = ctx.Bind(req); nil != err {
		return resp.ToJSON(ctx)
	}

	if err = ctx.Validate(req); nil != err {
		return resp.ToJSON(ctx)
	}

	if resp.Data, err = handler.api.Service.ForgotPassword(ctx, *req); nil != err {
		return resp.ToJSON(ctx)
	}

	return resp.ToJSON(ctx)
}

func (handler *userHandler) ChangePassword(ctx echo.Context) error {
	var req = new(request.UserChangePassword)
	var resp = response.New()
	var err error

	if err = ctx.Bind(req); nil != err {
		return resp.ToJSON(ctx)
	}

	if err = ctx.Validate(req); nil != err {
		return resp.ToJSON(ctx)
	}

	if resp.Data, err = handler.api.Service.ChangePassword(ctx, *req); nil != err {
		return resp.ToJSON(ctx)
	}

	return resp.ToJSON(ctx)
}
