package response

import (
	"time"
)

type User struct {
	ID             string  `json:"id"`
	FullName       string  `json:"full_name"`
	Username       string  `json:"username"`
	Email          string  `json:"email"`
	PhoneNumber    string  `json:"phone_number"`
	Balance        float64 `json:"balance"`
	OauthToken     string  `json:"oauth_token"`
	RefreshToken   string  `json:"refresh_token"`
	Token          string  `json:"token"`
	ProfilePicture string  `json:"profile_picture"`
	Status         int     `json:"status"`
	Type           int     `json:"type"`
}

type UserLogin struct {
	User
	LoginInfo UserLoginInfo `json:"login_info"`
}

type UserLoginInfo struct {
	Success bool `json:"success"`
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

type UserSendVerificationCode struct {
	Email     string    `json:"email"`
	Success   bool      `json:"success"`
	ExpiredAt time.Time `json:"expired_at"`
	Token     string    `json:"token"`
}

type UserVerification struct {
	User
	Success bool `json:"success"`
}
