package service

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/Simo672K/issue-tracker/utils"
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

type EmailContent struct {
	Subject string
	Content string
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

func (smtpc *SendgridSMTP) Init(host, port, user, password string) *EmailServConfig {
	smtpc.conf.host = host
	smtpc.conf.port = port
	smtpc.conf.user = user
	smtpc.conf.password = password
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

func EmailService(recievers []string, content *EmailContent) error {
	if err := LoadSMTPConfig().Init(
		SMTPHost,
		SMTPPort,
		SMTPUsername,
		SMTPPassword,
	).New(
		"<sender-email>",
		recievers,
		content.Subject,
		content.Content).Send(); err != nil {
		return err
	}
	return nil
}

func NewMockUserEmailVerification() string {
	userID := utils.StrUniqueId()
	verifId := CreateVerification(userID)

	return verifId
}

func SendVerificationEmail() {
	vid := NewMockUserEmailVerification()
	content := fmt.Sprintf("Please verify your account, this is your verification link: http://localhost:3000/api/v1/verify-email/%s", vid)
	c := &EmailContent{
		Subject: "Verify your email",
		Content: content,
	}

	EmailService([]string{"<reciever-email>"}, c)
}
