package builder

import (
	"github.com/google/uuid"
	"github.com/purwalenta/purwalenta/pkg/config"
	"github.com/purwalenta/purwalenta/pkg/entity"
	"github.com/purwalenta/purwalenta/pkg/service/request"
	"github.com/purwalenta/purwalenta/pkg/util"
)

func UserSignUpVerificationEmailTemplate(uuid uuid.UUID, req request.UserSignUp) entity.TemplateEmail {
	cfg := config.GetConfig()

	accessToken, _ := util.GenerateUserLoginToken(cfg, entity.User{
		ID:          uuid.String(),
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	})

	return entity.TemplateEmail{
		EmailBase: entity.EmailBase{
			From:    cfg.App.SignUpEmailAgent,
			To:      req.Email,
			Subject: "User Email Verification",
		},
		TemplateFlag: entity.SignUpEmailTemplateFlag,
		StringReplacers: map[string]string{
			"{{UserFullName}}":      req.FullName,
			"{{RegistrationToken}}": accessToken,
		},
	}
}
