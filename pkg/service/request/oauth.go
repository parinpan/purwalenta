package request

type OauthExchange struct {
	Source string `json:"source" form:"source" validate:"required"`
	Code   string `json:"code" form:"code" validate:"required"`
}
