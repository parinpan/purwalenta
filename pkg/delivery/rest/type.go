package rest

import (
	"github.com/labstack/echo"
)

type Configuration struct {
	Address string
}

type Response struct {
	Header   interface{} `json:"header"`
	Messages []string    `json:"messages"`
	Success  bool        `json:"success"`
	Data     interface{} `json:"data"`
}

func NewResponse(ctx echo.Context) *Response {
	return &Response{
		Header:   ctx.Request().Header,
		Messages: make([]string, 0),
		Success:  false,
		Data:     nil,
	}
}
