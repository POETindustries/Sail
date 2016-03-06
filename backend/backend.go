package backend

import (
	"bytes"
	"net/http"
	"sail/pages"
	"sail/session"
	"sail/tmpl"
	"sail/users"
)

// LoginPage returns the login page, asking for user credentials.
//
// On a fresh installation, the root user is 'admin' and the password
// is 'password'. Leave it like this at your own peril.
func LoginPage(req *http.Request) (*bytes.Buffer, *http.Cookie) {
	cookie, _ := req.Cookie("session")
	if cookie != nil && session.DB().Has(cookie.Value) {
		session.DB().Start(cookie.Value)
		return office(req), nil
	}
	if ok, msg := loginConfirm(req); !ok {
		return loginPage(msg), nil
	}
	s := session.New(req, req.PostFormValue("user"))
	session.DB().Add(s)
	c := http.Cookie{Name: "session", Value: s.ID}

	return office(req), &c
}

func loginConfirm(req *http.Request) (bool, string) {
	u := req.PostFormValue("user")
	p := req.PostFormValue("pass")
	if u != "" && p != "" {
		if users.Verify(u, p) {
			return true, ""
		}
		return false, "Wrong login credentials!"
	}
	return false, ""
}

func loginPage(msg string) *bytes.Buffer {
	presenter := pages.NewFromCache("/login")
	if msg != "" {
		presenter.Message = msg
		presenter.HasMessage = true
	}
	if markup, err := presenter.Compile(); err == nil {
		return markup
	}
	return bytes.NewBufferString(tmpl.NOTFOUND404)
}

func office(req *http.Request) *bytes.Buffer {
	println(req.URL.RequestURI())
	presenter := pages.NewWithURL(req.URL.RequestURI(), false)
	if markup, err := presenter.Compile(); err == nil {
		return markup
	}
	return bytes.NewBufferString(tmpl.NOTFOUND404)
}
