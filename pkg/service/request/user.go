package request

type UserLogin struct {
	ID       string `json:"id" form:"id" query:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

type UserSignUp struct {
}
