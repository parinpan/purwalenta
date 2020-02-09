package driver

import (
	"github.com/purwalenta/purwalenta/pkg/entity"
)

var (
	templateMailHandler = map[EmailDriver]templateMailFn{
		GoMailDriver: sendTemplateMailGoMailDriver,
	}
)

func SendTemplateMail(driver EmailDriver, email entity.TemplateEmail) error {
	if fn, ok := templateMailHandler[driver]; ok {
		return fn(email)
	}

	return sendTemplateMailGoMailDriver(email)
}
