package main

import (
	"bytes"
	"io"
	"net/http"
	"sail/conf"
	"sail/dbase"
	"sail/page"
)

// FrontendHandler handles all requests that are coming from site visitors.
// It parses the request url and calls the functions necessary for generating
// a valid page that is send to the client.
func frontendHandler(writer http.ResponseWriter, req *http.Request) {
	var p *page.Page
	var b bytes.Buffer

	if db := dbase.Open("sl_main"); db == nil {
		p = page.Load404()
	} else {
		p = page.Builder("home", db)
	}
	if err := p.Frame.Execute(&b, p); err != nil {
		println(err.Error())
		io.WriteString(writer, page.NOTFOUND404)
	} else {
		b.WriteTo(writer)
	}
}

// BackendHandler handles connections to the administrative interface.
func backendHandler(writer http.ResponseWriter, req *http.Request) {
	// TODO check for session cookie, show login page if not present
}

func main() {
	conf.InitConf()
	http.HandleFunc("/", frontendHandler)
	http.Handle("/img/", http.FileServer(http.Dir(conf.CWD)))
	http.Handle("/js/", http.FileServer(http.Dir(conf.CWD)))
	http.Handle("/theme/", http.FileServer(http.Dir(conf.CWD)))
	http.HandleFunc("/office/", backendHandler)
	http.ListenAndServe(":8080", nil)
}
