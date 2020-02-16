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
	Email string `json:"email" form:"email" validate:"required"`
}

type UserVerification struct {
	Email string `json:"email" form:"email" validate:"required"`
	Code  int    `json:"code" form:"code" validate:"required"`
}

type UserForgotPassword struct {
	Email string `json:"email" form:"email" validate:"required"`
}

type UserChangePassword struct {
	Email       string `json:"email" form:"email" validate:"required"`
	OldPassword string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
}
