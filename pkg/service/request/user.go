package request

type UserLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

type UserSignUp struct {
	Username    string `json:"username" form:"username"`
	FullName    string `json:"full_name" form:"full_name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Password    string `json:"password" form:"password"`
	Type        int    `json:"user_type" form:"type"`
}
