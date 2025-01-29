package smtp

import (
	"crypto/tls"
	"errors"
	"net/smtp"
)

type Smtp struct {
	Host       string
	Port       string
	User       string
	Password   string
	TLSEnabled bool
	AuthMethod string
}

func NewSmtp(host, port, user, password string, tlsEnabled bool, authMethod string) *Smtp {
	return &Smtp{
		Host:       host,
		Port:       port,
		User:       user,
		Password:   password,
		TLSEnabled: tlsEnabled,
		AuthMethod: authMethod,
	}
}

func (s *Smtp) Send(email Email) error {
	auth, err := s.buildAuth()
	if err != nil {
		return err
	}

	email.buildData()
	msg, err := email.buildMsgFromTemplate()
	if err != nil {
		return err
	}

	if s.TLSEnabled {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         s.Host,
		}

		conn, err := tls.Dial("tcp", s.Host+":"+s.Port, tlsConfig)
		if err != nil {
			return err
		}

		client, err := smtp.NewClient(conn, s.Host)
		if err != nil {
			return err
		}

		defer client.Close()
		if err := client.Auth(auth); err != nil {
			return err
		}

		if err := sendEmail(client, email, msg); err != nil {
			return err
		}
	} else {
		err := smtp.SendMail(s.Host+":"+s.Port, auth, email.From, email.createRecipientList(), []byte(msg))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Smtp) buildAuth() (smtp.Auth, error) {
	if s.AuthMethod == "PLAIN" {
		return smtp.PlainAuth("", s.User, s.Password, s.Host), nil
	}
	return nil, errors.New("unsupported auth method")
}

func sendEmail(client *smtp.Client, email Email, msg string) error {
	if err := client.Mail(email.From); err != nil {
		return err
	}

	for _, addr := range email.createRecipientList() {
		if err := client.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	return client.Quit()
}
