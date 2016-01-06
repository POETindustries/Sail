package backend

import (
	"bytes"
	"net/http"
	"sail/pages"
	"sail/session"
	"sail/tmpl"
)

// LoginPage returns the login page, asking for user credentials
func LoginPage(req *http.Request) (*bytes.Buffer, *http.Cookie) {
	cookie, _ := req.Cookie("session")
	if cookie != nil && session.DBInstance().Has(cookie.Value) {
		return bytes.NewBufferString("All well, session found"), nil
	}
	if ok, msg := loginConfirm(req); !ok {
		return loginPage(msg), nil
	}
	s := session.New(req, "whutman")
	session.DBInstance().Add(s)
	c := http.Cookie{Name: "session", Value: s.ID}

	return bytes.NewBufferString("All well, session created"), &c
}

func loginConfirm(req *http.Request) (bool, string) {
	user := req.PostFormValue("user")
	pass := req.PostFormValue("pass")
	if user != "" && pass != "" {
		return true, ""
	}
	return false, ""
}

func loginPage(msg string) *bytes.Buffer {
	presenter := pages.NewWithURL("/login")
	if msg != "" {
		presenter.Message = msg
		presenter.HasMessage = true
	}
	if markup, err := presenter.Compile(); err == nil {
		return markup
	}
	return bytes.NewBufferString(tmpl.NOTFOUND404)
}
