package driver

import "github.com/parinpan/purwalenta/pkg/entity"

type (
	EmailDriver int

	basicMailFn    func(entity.BasicEmail) error
	templateMailFn func(entity.TemplateEmail) error
)

const (
	emailTemplateExt  = ".html"
	emailTemplatePath = "/var/email/template"
	GoMailDriver      = 1
)
