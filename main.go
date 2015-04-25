package main

import (
	_ "html/template"
	"net/http"
	_ "sail/conf"
	"sail/dbase"
	"sail/page"
)

func frontendHandler(writer http.ResponseWriter, req *http.Request) {
	db := dbase.Open("sl_main")
	p := page.Builder("home", db)
	p.Frame.Execute(writer, p)
}

func backendHandler(writer http.ResponseWriter, req *http.Request) {
	// check for session cookie, show login page if not present
}

func main() {
	http.HandleFunc("/", frontendHandler)
	http.HandleFunc("/office/", backendHandler)
	http.ListenAndServe(":8080", nil)
}

func testCurrent() {
}
