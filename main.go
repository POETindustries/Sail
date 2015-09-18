package main

import (
	"bytes"
	"io"
	"net/http"
	"sail/core/conf"
	"sail/core/dbase"
	"sail/core/page"
	"sail/widget"
)

var conn *dbase.Conn

func main() {
	config := conf.Instance()
	initPlugins()

	if conn = dbase.New(); conn != nil {
		http.HandleFunc("/", frontendHandler)
		http.HandleFunc("/office/", backendHandler)

		http.Handle("/img/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/js/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/theme/", http.FileServer(http.Dir(config.Cwd)))

		http.ListenAndServe(":8080", nil)
	}
}

// FrontendHandler handles all requests that are coming from site visitors.
func frontendHandler(writer http.ResponseWriter, req *http.Request) {
	var b bytes.Buffer
	p := page.New(req.URL.RequestURI(), conn)

	if conn.Verify() && p.Execute(&b) == nil {
		b.WriteTo(writer)
	} else {
		io.WriteString(writer, page.NOTFOUND404)
	}
}

// BackendHandler handles connections to the administrative interface.
func backendHandler(writer http.ResponseWriter, req *http.Request) {
	// TODO check for session cookie, show login page if not present
}

func initPlugins() {
	widget.Init()
}
