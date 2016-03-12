package main

import (
	"fmt"
	"net/http"
	"sail/conf"
	"sail/response"
	"sail/storage"
	"sail/user"
	"sail/user/session"
	"time"
)

func main() {
	config := conf.Instance()
	if psqldb.Instance() != nil {
		storage.ExecCreateInstructs()
		http.HandleFunc("/", frontendHandler)
		http.HandleFunc("/office/", backendHandler)
		http.HandleFunc("/login", loginHandler)
		http.HandleFunc("/favicon.ico", iconHandler)
		http.Handle("/files/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/js/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/theme/", http.FileServer(http.Dir(config.Cwd)))
		http.ListenAndServe(":8080", nil)
	}
}

func frontendHandler(wr http.ResponseWriter, req *http.Request) {
	t1 := time.Now().Nanosecond()

	if psqldb.Instance().Verify() {
		response.New(wr, req).Serve()
	}

	t2 := time.Now().Nanosecond()
	fmt.Printf("Time to serve page: %d microseconds\n", (t2-t1)/1000)
}

func iconHandler(wr http.ResponseWriter, req *http.Request) {
	http.ServeFile(wr, req, conf.Instance().Cwd+"/favicon.ico")
}

func backendHandler(wr http.ResponseWriter, req *http.Request) {
	if psqldb.Instance().Verify() {
		cookie, _ := req.Cookie("session")
		if cookie != nil && session.DB().Has(cookie.Value) {
			session.DB().Start(cookie.Value)
			r := response.New(wr, req)
			r.FallbackURL = "/office/home"
			r.Serve()
		} else {
			loginHandler(wr, req)
		}
	}
}

func loginHandler(wr http.ResponseWriter, req *http.Request) {
	u := req.PostFormValue("user")
	p := req.PostFormValue("pass")
	r := response.New(wr, req)
	if u != "" && p != "" {
		if user.Verify(u, p) {
			s := session.New(req, req.PostFormValue("user"))
			session.DB().Add(s)
			c := http.Cookie{Name: "session", Value: s.ID}
			http.SetCookie(wr, &c)
			r.FallbackURL = "/office/home"
			r.Serve()
			return
		}
		r.Message = "Wrong login credentials!"
	}
	r.URL = "/logn"
	r.Serve()
}
