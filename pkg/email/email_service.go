package email

import (
	"log"
	"net/smtp"

	"github.com/echewisi/numeris_assessment/pkg/config"
)

type EmailService struct {
	host     string
	port     string
	username string
	password string
	from     string
}

// NewEmailService initializes the EmailService with configurations
func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		host:     cfg.Email.SMTPHost,
		port:     cfg.Email.SMTPPort,
		username: cfg.Email.Username,
		password: cfg.Email.Password,
		from:     cfg.Email.From,
	}
}

// SendEmail sends an email to a specific recipient
func (s *EmailService) SendEmail(to, subject, body string) error {
	// Set up authentication information
	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	// Construct the message
	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body + "\r\n")

	// Send email
	err := smtp.SendMail(s.host+":"+s.port, auth, s.from, []string{to}, message)
	if err != nil {
		log.Printf("Failed to send email to %s: %v", to, err)
		return err
	}

	log.Printf("Email sent successfully to %s", to)
	return nil
}
