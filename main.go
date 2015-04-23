package main

import (
	"html/template"
	"net/http"
	"sail/conf"
)

func frontendHandler(writer http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles(conf.DOCROOT + "index.html")
	if err != nil {
		println(err.Error())
	} else {
		t.Execute(writer, nil)
	}
}

func backendHandler(writer http.ResponseWriter, req *http.Request) {
	// check for session cookie, show login page if not present
}

func main() {
	http.HandleFunc("/", frontendHandler)
	http.HandleFunc("/office/", backendHandler)
	http.ListenAndServe(":8080", nil)
}
