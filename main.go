package main

import (
	"bytes"
	"io"
	"net/http"
	"sail/dbase"
	"sail/page"
)

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

func backendHandler(writer http.ResponseWriter, req *http.Request) {
	// check for session cookie, show login page if not present
}

func main() {

	http.HandleFunc("/", frontendHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.HandleFunc("/office/", backendHandler)
	http.ListenAndServe(":8080", nil)
}

func testCurrent() {
}
