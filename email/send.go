package email

import (
	"errors"
	"net/smtp"

	"github.com/darkoatanasovski/htmltags"
	"github.com/jordan-wright/email"
	"log/slog"
)

// Config holds the email configuration
type Config struct {
	// Host is the SMTP server host
	Host string

	// Port is the SMTP server port
	Port string

	// Username is the SMTP server username
	Username string

	// Password is the SMTP server password
	Password string

	// Logger for logging errors
	Logger *slog.Logger
}

// SendOptions defines the options for sending an email
type SendOptions struct {
	// From is the email sender
	From string

	// FromName is the name of the sender (unused for now)
	FromName string

	// To is the list of recipients
	To []string

	// Bcc is the list of BCC recipients
	Bcc []string

	// Cc is the list of CC recipients
	Cc []string

	// Subject is the email subject
	Subject string

	// HtmlBody is the HTML content of the email
	HtmlBody string

	// TextBody is the plain text content of the email
	TextBody string
}

// Sender defines the interface for sending emails
type Sender interface {
	// Send sends an email with the given options
	Send(options SendOptions) error
}

// SMTPSender implements the Sender interface using SMTP
type SMTPSender struct {
	config Config
}

// NewSMTPSender creates a new SMTP email sender
func NewSMTPSender(config Config) Sender {
	return &SMTPSender{
		config: config,
	}
}

// Send sends an email using SMTP
func (s *SMTPSender) Send(options SendOptions) error {
	host := s.config.Host
	port := s.config.Port
	user := s.config.Username
	pass := s.config.Password
	addr := host + ":" + port

	if options.From == "" {
		return errors.New("from is required")
	}

	if len(options.To) == 0 {
		return errors.New("to is required")
	}

	if options.Subject == "" {
		return errors.New("subject is required")
	}

	if options.HtmlBody == "" {
		return errors.New("html is required")
	}

	if options.TextBody == "" {
		nodes, errStripped := htmltags.Strip(options.HtmlBody, []string{}, true)

		if errStripped == nil {
			options.TextBody = nodes.ToString() // returns stripped HTML string
		}
	}

	e := email.NewEmail()
	e.From = options.From
	e.To = options.To
	e.Bcc = options.Bcc
	e.Cc = options.Cc
	e.Subject = options.Subject
	e.Text = []byte(options.TextBody)
	e.HTML = []byte(options.HtmlBody)
	
	var auth smtp.Auth
	if user == "" {
		auth = nil
	} else {
		auth = smtp.PlainAuth("", user, pass, host)
	}

	err := e.Send(addr, auth)

	if err != nil && s.config.Logger != nil {
		s.config.Logger.Error("Error sending email", "error", err.Error())
		return err
	}

	return nil
}
