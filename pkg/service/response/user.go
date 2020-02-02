package response

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UserLogin struct {
	User
}

type UserSignUp struct {
	User
}
