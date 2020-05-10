package response

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/parinpan/purwalenta/pkg/errord"
)

type Configuration struct {
	Address     string
	IdleTimeout time.Duration
	ReadTimeout time.Duration
}

type response struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

func New() *response {
	return new(response)
}

func (resp *response) ToJSON(ctx echo.Context) error {
	errorComponent, _ := ctx.Get("errord_error").(errord.ErrorComponent)
	httpStatusCode := http.StatusOK

	if errorComponent.HttpStatus > 0 {
		httpStatusCode = errorComponent.HttpStatus
	} else {
		errorComponent.HttpStatus = httpStatusCode
	}

	// set final error component to response
	resp.Error = errorComponent

	return ctx.JSON(httpStatusCode, resp)
}
