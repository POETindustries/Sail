package main

import (
	"bytes"
	"io"
	"net/http"
	"sail/conf"
	"sail/dbase"
	"sail/page"
)

var config *conf.Config
var conn *dbase.Conn

// FrontendHandler handles all requests that are coming from site visitors.
// It parses the request url and calls the functions necessary for generating
// a valid page that is send to the client.
func frontendHandler(writer http.ResponseWriter, req *http.Request) {
	var p *page.Page
	var b bytes.Buffer

	if !conn.Verify() {
		p = page.Load404()
	} else {
		p = page.Build("home", conn.DB)
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
	config := conf.New()
	if conn = dbase.New(config); conn != nil {
		http.HandleFunc("/", frontendHandler)
		http.HandleFunc("/office/", backendHandler)

		http.Handle("/img/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/js/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/theme/", http.FileServer(http.Dir(config.Cwd)))

		http.ListenAndServe(":8080", nil)
	}
}
