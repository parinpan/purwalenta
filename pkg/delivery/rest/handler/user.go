package handler

import (
	"net/http"

	"github.com/labstack/echo"
	apiPkg "github.com/purwalenta/purwalenta/pkg/api"
	"github.com/purwalenta/purwalenta/pkg/delivery/rest/response"
	"github.com/purwalenta/purwalenta/pkg/service/request"
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
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err = ctx.Validate(req); nil != err {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if resp.Data, err = handler.api.Service.Login(ctx, *req); nil != err {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (handler *userHandler) SignUp(ctx echo.Context) error {
	var req = new(request.UserSignUp)
	var resp = response.New()
	var err error

	if err = ctx.Bind(req); nil != err {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err = ctx.Validate(req); nil != err {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if resp.Data, err = handler.api.Service.SignUp(ctx, *req); nil != err {
		ctx.Logger().Error(err)
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	return ctx.JSON(http.StatusOK, resp)
}