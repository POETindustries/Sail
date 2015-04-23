package main

import (
	"html/template"
	"net/http"
	"sail/conf"
	"sail/dbase"
)

var db = dbase.Open("hi")

func reqHandler(writer http.ResponseWriter, req *http.Request) {
	row := db.QueryRow("select language from sl_page_meta where domain=?", "default")
	var lang string
	err1 := row.Scan(&lang)
	t, err := template.ParseFiles(conf.DOCROOT + "index.html")
	if err != nil {
		println(err.Error())
	} else {
		if err1 != nil {
			println(err1.Error())
		}
		t.Execute(writer, nil)
	}
}

func main() {
	http.HandleFunc("/", reqHandler)
	if db != nil {
		println("db connected")
	}
	http.ListenAndServe(":8080", nil)
}
