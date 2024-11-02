package service

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/lpernett/godotenv"
)

type SendgridSMTP struct {
	conf EmailServConfig
}

type EmailServConfig struct {
	host     string
	port     string
	user     string
	password string
}

type Email struct {
	from    string
	to      []string
	subject string
	message string
	auth    smtp.Auth
	addr    string
}

var (
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("failed to load .env file")
	}
	SMTPHost = os.Getenv("SMTP_HOST")
	SMTPPort = os.Getenv("SMTP_PORT")
	SMTPPassword = os.Getenv("SMTP_PASSWORD")
	SMTPUsername = os.Getenv("SMTP_USERNAME")
}

func LoadSMTPConfig() *SendgridSMTP {
	return &SendgridSMTP{
		conf: EmailServConfig{}}
}

func (smtpc *SendgridSMTP) Init() *EmailServConfig {
	smtpc.conf.host = SMTPHost
	smtpc.conf.port = SMTPPort
	smtpc.conf.password = SMTPPassword
	smtpc.conf.user = SMTPUsername
	return &smtpc.conf
}

func (esc *EmailServConfig) New(from string, to []string, subject string, msg string) *Email {
	return &Email{
		from:    from,
		to:      to,
		message: msg,
		subject: subject,
		auth:    smtp.PlainAuth("", esc.user, esc.password, esc.host),
		addr:    esc.host + ":" + esc.port,
	}
}

func (e *Email) Send() error {
	emailBody := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		e.from, e.to[0], e.subject, e.message)
	if err := smtp.SendMail(e.addr, e.auth, e.from, e.to, []byte(emailBody)); err != nil {
		return err
	}

	return nil
}

func EmailService() error {
	if err := LoadSMTPConfig().Init().New(
		"<sender>",
		[]string{"<reciever>"},
		"test email",
		"test email content").Send(); err != nil {
		return err
	}
	return nil
}
