package email

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"mime"
	"net/http"
	"net/smtp"
	"path/filepath"
	"sail/conf"
	"strconv"
	"strings"
	"text/template"
)

const (
	msgPlain = "From: {{if .From.Name}}\"{{.From.Name}}\" <{{.From.Address}}>{{else}}{{.From.Address}}{{end}}\r\n" +
		"To: {{range .To}}{{if .Name}}\"{{.Name}}\" <{{.Address}}>{{else}}{{.Address}}{{end}},{{end}}\r\n" +
		"Subject: {{.Subject}}\r\nContent-Type: text/plain; charset=utf-8\r\nContent-Transfer-Encoding: 8bit\r\n\r\n" +
		"{{.Body}}"

	msgHTML = "From: {{if .From.Name}}\"{{.From.Name}}\" <{{.From.Address}}>{{else}}{{.From.Address}}{{end}}\r\n" +
		"To: {{range .To}}{{if .Name}}\"{{.Name}}\" <{{.Address}}>{{else}}{{.Address}}{{end}},{{end}}\r\n" +
		"Subject: {{.Subject}}\r\nContent-Type: text/html; charset=utf-8\r\nContent-Transfer-Encoding: 8bit\r\n\r\n" +
		"<div id='msg'>{{.Body}}</div>"

	msgMIME = "From: {{if .From.Name}}\"{{.From.Name}}\" <{{.From.Address}}>{{else}}{{.From.Address}}{{end}}\r\n" +
		"To: {{range .To}}{{if .Name}}\"{{.Name}}\" <{{.Address}}>{{else}}{{.Address}}{{end}},{{end}}\r\n" +
		"Subject: {{.Subject}}\r\nMIME-Version: 1.0\r\nContent-Type: multipart/mixed; boundary=\"---uuddlrlrba\"\r\n\r\n" +
		"-----uuddlrlrba\r\nContent-Type: text/plain; charset=utf-8\r\nContent-Transfer-Encoding: 8bit\r\n\r\n{{.Body}}\r\n" +
		"{{range .Files}}\r\n-----uuddlrlrba\r\nContent-Type: {{.Type}}\"\r\nContent-Transfer-Encoding: {{.Encoding}}\r\n" +
		"Content-Disposition: \"attachment\"; filename=\"{{.Name}}\"\r\n\r\n{{.Data}}\r\n{{end}}-----uuddlrlrba--\r\n"
)

var (
	tmplPlain, _ = template.New("msg").Parse(msgPlain)
	tmplHTML, _  = template.New("msg").Parse(msgHTML)
	tmplMIME, _  = template.New("msg").Parse(msgMIME)
)

// Sender represents the sending party of email traffic. The
// Sender stores identity information and login credentials
// for connecting to an SMTP server.
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

// Recipient represents the receiving party of email traffic.
// In order to be able to receive emails, a valid address
// must be provided, along with an optional name.
type Recipient struct {
	Name    string
	Address string
}

// Email wraps package smtp functionality and provides an
// easy way to send an email to multiple recipients.
type Email struct {
	From     *Sender
	To       []Recipient
	Subject  string
	Body     string
	Template *template.Template

	files []*File
}

// New returns a new Email, ready to be filled with content
// and initialized with the user argument or the system-wide
// email user if user was nil.
func New(sender *Sender) *Email {
	s := systemSender()
	if sender != nil {
		s = sender
	}
	if s.Auth == nil {
		s.ParseAuth()
	}
	return &Email{
		From:     s,
		To:       []Recipient{},
		Template: tmplPlain,
		files:    []*File{}}
}

// NewHTML returns a new HTML-enabled Email, initialized
// with the user argument or the system-wide email user if
// user was nil.
func NewHTML(sender *Sender) *Email {
	e := New(sender)
	e.Template = tmplHTML
	return e
}

// NewMultipart returns a new multipart Email, initialized
// with the user argument or the system-wide email user if
// user was nil. This type of email is used for messages
// with attachments.
func NewMultipart(sender *Sender) *Email {
	e := New(sender)
	e.Template = tmplMIME
	return e
}

// AddFile adds a new file to the list of attachments. An
// email can have an arbitrary number of attachments sent
// along.
//
// If the file could be read, a proper MIME type is inferred
// and the data is encoded and formatted for compliance with
// RFC 2046 and RFC 2183.
func (e *Email) AddFile(name string) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}
	f := &File{Encoding: "base64", Name: filepath.Base(name)}
	mByExt := mime.TypeByExtension(filepath.Ext(name))
	mByHead := http.DetectContentType(data)
	if mByExt == "" {
		f.Type = mByHead
	} else if strings.Split(mByExt, "/")[0] != strings.Split(mByHead, "/")[0] {
		f.Type = mByHead
	} else {
		f.Type = mByExt
	}
	for i, d := 0, base64.StdEncoding.EncodeToString(data); i < len(d); i += 72 {
		if i+72 < len(d) {
			f.Data += d[i:i+72] + "\r\n"
		} else {
			f.Data += d[i:]
		}
	}
	e.files = append(e.files, f)
}

// Files returns the list of attachments in a way suitable
// for email message composition. Its primary use lies in
// providing data for template execution.
func (e *Email) Files() []*File {
	return e.files
}

// Send transmits the Email and returns nil if successful.
// Otherwise, an error describing the problem is returned.
func (e *Email) Send() error {
	if len(e.To) == 0 {
		return &ErrNoRCPT{}
	}
	var msg bytes.Buffer
	if err := e.Template.Execute(&msg, e); err != nil {
		return err
	}
	host := e.From.Host + ":" + strconv.FormatUint(uint64(e.From.Port), 10)
	to := []string{}
	for _, r := range e.To {
		to = append(to, r.Address)
	}
	return smtp.SendMail(host, e.From.Auth, e.From.Address, to, msg.Bytes())
}

// systemSender returns the application's default mail user.
// Mails sent with this user should be considered as mails
// coming from "the app".
func systemSender() *Sender {
	return &Sender{
		Name:    conf.Instance().MailUser,
		Address: conf.Instance().MailAddress,
		Pass:    conf.Instance().MailPass,
		Host:    conf.Instance().MailHostSMTP,
		Port:    conf.Instance().MailPortSMTP}
}

// File wraps an email attachment and exposes MIME data.
// Files should be added to emails only to the existing
// methods on the Email type in order to insure correct
// formatting and encoding.
type File struct {
	Type     string
	Encoding string
	Data     string
	Name     string
}
