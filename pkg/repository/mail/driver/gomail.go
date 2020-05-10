package driver

import (
	"crypto/tls"

	"github.com/parinpan/purwalenta/pkg/config"
	"github.com/parinpan/purwalenta/pkg/entity"
	"gopkg.in/gomail.v2"
)

func sendTemplateMailGoMailDriver(email entity.TemplateEmail) error {
	cfg := config.GetConfig().SMTP
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", email.From)
	mailer.SetHeader("To", email.To)
	mailer.SetHeader("Subject", email.Subject)
	mailer.SetBody("text/html", buildTemplateEmailBody(email))

	dialer := gomail.NewPlainDialer(cfg.Outgoing.Server, cfg.Outgoing.Port, cfg.Username, cfg.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return dialer.DialAndSend(mailer)
}
