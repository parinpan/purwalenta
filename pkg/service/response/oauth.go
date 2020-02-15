package response

type OauthExchange struct {
	User
	Token   string `json:"token"`
	Success bool   `json:"success"`
}
