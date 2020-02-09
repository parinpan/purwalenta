package driver

import (
	"io/ioutil"
	"strings"

	"github.com/purwalenta/purwalenta/pkg/entity"
)

func buildTemplateEmailBody(email entity.TemplateEmail) string {
	var strReplacers []string

	for oldStr, newStr := range email.StringReplacers {
		strReplacers = append(strReplacers, oldStr, newStr)
	}

	fileBytes, _ := ioutil.ReadFile(emailTemplatePath + "/" + string(email.TemplateFlag) + emailTemplateExt)
	replacedTemplate := strings.NewReplacer(strReplacers...).Replace(string(fileBytes))

	return replacedTemplate
}
