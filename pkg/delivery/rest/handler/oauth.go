package handler

import (
	"net/http"

	"github.com/labstack/echo"
	apiPkg "github.com/parinpan/purwalenta/pkg/api"
	"github.com/parinpan/purwalenta/pkg/delivery/rest/response"
	"github.com/parinpan/purwalenta/pkg/service/request"
)

type oauthHandler struct {
	api *apiPkg.OauthAPI
}

func NewOauthHandler() *oauthHandler {
	return &oauthHandler{
		api: apiPkg.NewOauthAPI(apiPkg.DefaultOauthAPIFlag),
	}
}

func (handler *oauthHandler) Exchange(ctx echo.Context) error {
	var req = new(request.OauthExchange)
	var resp = response.New()
	var err error

	if err = ctx.Bind(req); nil != err {
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if err = ctx.Validate(req); nil != err {
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	if resp.Data, err = handler.api.Service.Exchange(ctx, *req); nil != err {
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	return ctx.JSON(http.StatusOK, resp)
}
