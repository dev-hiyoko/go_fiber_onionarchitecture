package configs

import (
	"hiyoko-fiber/pkg/mail/smtp"
	"hiyoko-fiber/utils"
)

func NewSmtpConf() (conf *smtp.Smtp) {
	smtpConf := smtp.NewSmtp(
		utils.Env("MAIL_HOST").GetString("localhost"),
		utils.Env("MAIL_PORT").GetString("1025"),
		utils.Env("MAIL_USERNAME").GetString(""),
		utils.Env("MAIL_PASSWORD").GetString(""),
		utils.Env("MAIL_TLS").GetBool(false),
		"PLAIN",
	)
	return smtpConf
}

func NewEmailConf() (*smtp.Email, error) {
	emailConf := smtp.NewEmail(
		utils.Env("MAIL_FROM_ADDRESS").GetString("no-reply@example.com"),
		utils.Env("MAIL_FROM_NAME").GetString("ひよこ"),
	)

	smtp.SetTemplateDirectory("resources/templates/email/")
	smtp.SetCommonDirectoryName("common/")
	err := emailConf.SetAssetsPngImages("resources/assets/images")

	return emailConf, err
}
