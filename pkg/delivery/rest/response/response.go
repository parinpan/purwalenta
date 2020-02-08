package response

import (
	"time"
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

type errorResponse struct {
	Code    string `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func New() *response {
	return new(response)
}
