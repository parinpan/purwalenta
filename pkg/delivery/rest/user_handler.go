package rest

import (
	"net/http"

	"github.com/labstack/echo"
	apiPkg "github.com/purwalenta/purwalenta/pkg/api"
	"github.com/purwalenta/purwalenta/pkg/service/request"
)

type userHandler struct {
	api *apiPkg.UserAPI
}

func newUserHandler() *userHandler {
	return &userHandler{
		api: apiPkg.NewUserAPI(apiPkg.DefaultUserAPIFlag),
	}
}

func (handler *userHandler) Login(ctx echo.Context) error {
	var req = request.UserLogin{}
	var resp = NewResponse(ctx)
	var err error

	if err = ctx.Bind(req); nil != err {
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	if resp.Data, err = handler.api.Service.Login(ctx, req); nil != err {
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	return ctx.JSON(http.StatusOK, resp)
}
