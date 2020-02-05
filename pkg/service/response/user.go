package response

type User struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Token       string `json:"token"`
	Type        int    `json:"type"`
}

type UserLogin struct {
	User
}

type UserSignUp struct {
	User
	SignUpInfo UserSignUpInfo `json:"sign_up_info"`
}

type UserSignUpInfo struct {
	Success          bool     `json:"success"`
	UserAlreadyExist bool     `json:"user_already_exist"`
	TakenFields      []string `json:"taken_fields"`
}
