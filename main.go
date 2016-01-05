package main

import (
	"fmt"
	"net/http"
	"sail/conf"
	"sail/pages"
	"sail/storage"
	"sail/storage/psqldb"
	"time"
)

func main() {
	config := conf.Instance()
	if psqldb.Instance() != nil {
		storage.ExecCreateInstructs()
		http.HandleFunc("/", frontendHandler)
		http.HandleFunc("/office/", backendHandler)
		http.HandleFunc("/favicon.ico", iconHandler)
		http.Handle("/files/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/js/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/theme/", http.FileServer(http.Dir(config.Cwd)))
		http.ListenAndServe(":8080", nil)
	}
}

func frontendHandler(writer http.ResponseWriter, req *http.Request) {
	t1 := time.Now().Nanosecond()

	if psqldb.Instance().Verify() {
		uri := req.URL.Path
		_ = req.URL.Query() // TODO needs to be handled somewhere
		pages.Serve(uri).WriteTo(writer)
	}

	t2 := time.Now().Nanosecond()
	fmt.Printf("Time to serve page: %d microseconds\n", (t2-t1)/1000)
}

func iconHandler(writer http.ResponseWriter, req *http.Request) {
	http.ServeFile(writer, req, conf.Instance().Cwd+"/favicon.ico")
}

func backendHandler(writer http.ResponseWriter, req *http.Request) {
	// TODO check for session cookie, show login page if not present
}
