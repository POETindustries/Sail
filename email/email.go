package email

import (
	"bytes"
	"net/smtp"
	"strconv"
	"text/template"
)

// Sender stores login credentials for connecting to an smtp
// server.
type Sender struct {
	Name string

	Address string
	Pass    string
	Host    string
	Port    uint16

	Auth smtp.Auth
}

// ParseAuth composes the Sender's internal data and creates
// a valid authentication object for the mail server.
func (s *Sender) ParseAuth() {
	s.Auth = smtp.PlainAuth("", s.Address, s.Pass, s.Host)
}

// Email wraps package smtp functionality and provides an
// easy way to send an email to multiple recipients.
type Email struct {
	From     *Sender
	To       []string
	Subject  string
	Body     string
	Template *template.Template
}

// Send transmits the Email and returns nil if successful.
// Otherwise, an error describing the problem is returned.
func (e *Email) Send() error {
	host := e.From.Host + ":" + strconv.FormatUint(uint64(e.From.Port), 10)
	var msg bytes.Buffer
	if err := e.Template.Execute(&msg, e); err != nil {
		return err
	}
	return smtp.SendMail(host, e.From.Auth, e.From.Address, e.To, msg.Bytes())
}

// New returns a new Email, ready to be filled with content
// and initialized with the user argument or the system-wide
// email user if user was nil.
func New(sender *Sender) *Email {
	s := systemSender()
	if sender != nil {
		s = sender
	}
	return &Email{From: s}
}

func systemSender() *Sender {
	return &Sender{}
}
