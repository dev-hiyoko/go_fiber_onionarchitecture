package configs

import (
	"hiyoko-fiber/pkg/mail/smtp"
	"hiyoko-fiber/utils"
)

func NewSmtpConf() (conf smtp.Config) {
	conf.Host = utils.Env("MAIL_HOST").GetString("localhost")
	conf.Port = utils.Env("MAIL_PORT").GetString("1025")
	conf.User = utils.Env("MAIL_USERNAME").GetString("")
	conf.Password = utils.Env("MAIL_PASSWORD").GetString("")
	conf.TLSEnabled = utils.Env("MAIL_TLS").GetBool(false)
	conf.AuthMethod = "PLAIN"
	return
}

func NewEmailConf() (*smtp.Email, error) {
	email := &smtp.Email{
		From:     utils.Env("MAIL_FROM_ADDRESS").GetString("no-reply@example.com"),
		FromName: utils.Env("MAIL_FROM_NAME").GetString("ひよこ"),
	}
	err := email.SetAssetsPngImages("./assets/images")

	return email, err
}
