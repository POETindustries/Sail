package main

import (
	"bytes"
	"io"
	"net/http"
	"sail/conf"
	"sail/pages"
	"sail/storage"
	"sail/storage/psqldb"
	"sail/tmpl"
)

const docStart = "<!doctype html>"
const htmlOpen = "<html>"
const htmlClose = "</html>"

func main() {
	config := conf.Instance()
	if psqldb.Instance() != nil {
		storage.ExecCreateInstructs()
		http.HandleFunc("/", frontendHandler)
		http.HandleFunc("/office/", backendHandler)
		http.Handle("/img/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/js/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/theme/", http.FileServer(http.Dir(config.Cwd)))

		http.ListenAndServe(":8080", nil)
	}
}

func frontendHandler(writer http.ResponseWriter, req *http.Request) {
	b := bytes.NewBufferString(docStart + htmlOpen)
	p := pages.BuildWithURL(req.URL.RequestURI())
	if psqldb.Instance().Verify() && pages.Serve(p, b) == nil {
		b.WriteString(htmlClose)
		b.WriteTo(writer)
	} else {
		io.WriteString(writer, docStart+htmlOpen+tmpl.NOTFOUND404+htmlClose)
	}
}

func backendHandler(writer http.ResponseWriter, req *http.Request) {
	// TODO check for session cookie, show login page if not present
}
