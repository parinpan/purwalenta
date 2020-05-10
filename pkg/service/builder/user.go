package builder

import (
	"strconv"

	"github.com/parinpan/purwalenta/pkg/config"
	"github.com/parinpan/purwalenta/pkg/entity"
)

func UserSignUpVerificationEmailTemplate(verification entity.SignUpVerification) entity.TemplateEmail {
	cfg := config.GetConfig()

	return entity.TemplateEmail{
		EmailBase: entity.EmailBase{
			From:    cfg.App.SignUpEmailAgent,
			To:      verification.User.Email,
			Subject: "[Purwalenta] - User Email Verification",
		},
		TemplateFlag: entity.SignUpEmailTemplateFlag,
		StringReplacers: map[string]string{
			"{{UserFullName}}":     verification.User.FullName,
			"{{VerificationCode}}": strconv.Itoa(verification.VerificationCode),
		},
	}
}

func UserForgotPasswordTemplate(user entity.User, newPassword string) entity.TemplateEmail {
	cfg := config.GetConfig()

	return entity.TemplateEmail{
		EmailBase: entity.EmailBase{
			From:    cfg.App.SignUpEmailAgent,
			To:      user.Email,
			Subject: "[Purwalenta] - User Forgot Password",
		},
		TemplateFlag: entity.UserForgotPasswordTemplateFlag,
		StringReplacers: map[string]string{
			"{{UserFullName}}":    user.FullName,
			"{{UserNewPassword}}": newPassword,
		},
	}
}
