package main

import (
	"fmt"
	"net/http"
	"sail/backend"
	"sail/conf"
	"sail/frontend"
	"sail/object/cache"
	"sail/response"
	"sail/session"
	"sail/storage"
	"sail/user"
	"sail/user/group"
	"time"
)

var avgSrvTime = 0
var reqs = 0

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
	t := (t2 - t1)
	if t > 0 {
		reqs++
		avgSrvTime += t
		fmt.Printf("%3d Time to serve page: %d microseconds. Average: %d\n",
			reqs, t/1000, avgSrvTime/reqs/1000)
	}
}

func backendHandler(wr http.ResponseWriter, req *http.Request) {
	t1 := time.Now().Nanosecond()

	if storage.DB().Ping() == nil {
		cookie, _ := req.Cookie("id")
		if cookie != nil && session.DB().Has(cookie.Value) {
			s := session.DB().Get(cookie.Value)
			u := user.LoadNew(s.User)
			if b := group.NewBouncer(req); !b.Pass(u.ID()) {
				b.Sanitize("/office/")
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
	if usr, ok := session.Verify(user.New(u), p); ok {
		sess := session.New(req, req.PostFormValue("user"))
		session.DB().Add(sess)
		session.Users().Add(usr)
		c := http.Cookie{Name: "id", Value: sess.ID}
		http.SetCookie(wr, &c)
		if b := group.NewBouncer(req); !b.Pass(usr.ID()) {
			b.Sanitize("/office/")
		}
		r.URL = req.URL.Path
		r.Presenter = backend.New(sess, usr.(*user.User))
		sess.Start()
	} else {
		if u != "" || p != "" {
			r.Message = "Wrong login credentials!"
		}
		r.URL = "/office/login"
		r.Presenter = backend.New(nil, nil)
	}
	r.Serve()
}
