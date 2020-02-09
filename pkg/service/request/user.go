package request

type UserLogin struct {
	Username    string `json:"username" form:"username"`
	Email       string `json:"username" form:"email"`
	PhoneNumber string `json:"username" form:"phone_number"`
	Password    string `json:"password" form:"password" validate:"required"`
}

type UserSignUp struct {
	Username    string `json:"username" form:"username" validate:"required"`
	FullName    string `json:"full_name" form:"full_name"`
	Email       string `json:"email" form:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
	Type        int    `json:"user_type" form:"type" validate:"required"`
}

type UserSendVerificationCode struct {
	ID          string `json:"id" form:"id"`
	FullName    string `json:"full_name" form:"full_name"`
	Username    string `json:"username" form:"username"`
	Email       string `json:"email" form:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}
