package main

import (
	"fmt"
	"net/http"
	"sail/conf"
	"sail/page/backend"
	"sail/page/cache"
	"sail/page/frontend"
	"sail/response"
	"sail/storage"
	"sail/user"
	"sail/user/group"
	"sail/user/session"
	"time"
)

func main() {
	config := conf.Instance()
	if storage.DB() != nil {
		storage.ExecCreateInstructs()
		http.HandleFunc("/", frontendHandler)
		http.HandleFunc("/office/", backendHandler)
		http.Handle("/favicon.ico", http.FileServer(http.Dir(config.StaticDir)))
		http.Handle("/files/", http.FileServer(http.Dir(config.StaticDir)))
		http.Handle("/js/", http.FileServer(http.Dir(config.StaticDir)))
		http.Handle("/theme/", http.FileServer(http.Dir(config.StaticDir)))
		http.ListenAndServe(":8080", nil)
	}
}

func frontendHandler(wr http.ResponseWriter, req *http.Request) {
	t1 := time.Now().Nanosecond()

	if markup := cache.DB().Markup(req.URL.Path); markup != nil {
		wr.Write(markup)
	} else if storage.DB().Ping() == nil {
		r := response.New(wr, req)
		r.Presenter = frontend.New(r.Content(), r.Template())
		r.URL = r.Content().URL
		r.Serve()
	}

	t2 := time.Now().Nanosecond()
	fmt.Printf("Time to serve page: %d microseconds\n", (t2-t1)/1000)
}

func backendHandler(wr http.ResponseWriter, req *http.Request) {
	t1 := time.Now().Nanosecond()

	if storage.DB().Ping() == nil {
		cookie, _ := req.Cookie("session")
		if cookie != nil && session.DB().Has(cookie.Value) {
			s := session.DB().Get(cookie.Value)
			u := user.ByName(s.User)
			if b := group.NewBouncer(req); !b.Pass(u.ID) {
				b.Sanitize("/office/", "")
			}
			r := response.New(wr, req)
			r.FallbackURL = "/office/"
			r.Presenter = backend.New(s, u)
			s.Start()
			r.Serve()
		} else {
			loginHandler(wr, req)
		}
	}

	t2 := time.Now().Nanosecond()
	fmt.Printf("Time to serve page: %d microseconds\n", (t2-t1)/1000)
}

func loginHandler(wr http.ResponseWriter, req *http.Request) {
	u := req.PostFormValue("user")
	p := req.PostFormValue("pass")
	r := response.New(wr, req)
	if u != "" && p != "" {
		if user, ok := user.Verify(u, p); ok {
			s := session.New(req, req.PostFormValue("user"))
			session.DB().Add(s)
			c := http.Cookie{Name: "session", Value: s.ID}
			http.SetCookie(wr, &c)
			r.FallbackURL = "/office/"
			r.Presenter = backend.New(s, user)
			s.Start()
			r.Serve()
			return
		}
		r.Message = "Wrong login credentials!"
	}
	r.URL = "/office/login"
	r.Presenter = backend.New(nil, nil)
	r.Serve()
}
