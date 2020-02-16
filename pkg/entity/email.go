package entity

type (
	EmailTemplate string
)

const (
	SignUpEmailTemplateFlag        EmailTemplate = "user_registration"
	UserForgotPasswordTemplateFlag EmailTemplate = "user_forgot_password"
)

type EmailBase struct {
	From    string
	To      string
	CC      string
	Subject string
}

type BasicEmail struct {
	EmailBase
	Message string
}

type TemplateEmail struct {
	EmailBase
	TemplateFlag    EmailTemplate
	StringReplacers map[string]string
}
