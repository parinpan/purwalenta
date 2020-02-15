package response

type OauthExchange struct {
	User
	Token   string
	Success bool
}
