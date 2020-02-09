package builder

import (
	"strconv"

	"github.com/purwalenta/purwalenta/pkg/config"
	"github.com/purwalenta/purwalenta/pkg/entity"
)

func UserSignUpVerificationEmailTemplate(verification entity.SignUpVerification) entity.TemplateEmail {
	cfg := config.GetConfig()

	return entity.TemplateEmail{
		EmailBase: entity.EmailBase{
			From:    cfg.App.SignUpEmailAgent,
			To:      verification.User.Email,
			Subject: "User Email Verification",
		},
		TemplateFlag: entity.SignUpEmailTemplateFlag,
		StringReplacers: map[string]string{
			"{{UserFullName}}":     verification.User.FullName,
			"{{VerificationCode}}": strconv.Itoa(verification.VerificationCode),
		},
	}
}
