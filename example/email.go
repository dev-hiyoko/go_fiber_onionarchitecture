package main

import (
	"hiyoko-fiber/configs"
	"hiyoko-fiber/pkg/i18n"
)

func main() {
	config := configs.NewSmtpConf()
	email, _ := configs.NewEmailConf()
	email.To = []string{"test@test.com"}
	email.Subject = "Hiyoko"
	email.TemplateFileName = "example.tmpl"
	lang, _ := i18n.GetLanguageMap(i18n.Japanese)
	email.LangData = lang
	email.SetData(map[string]string{
		"Code": "123456",
	})
	_ = config.Send(*email)
}
