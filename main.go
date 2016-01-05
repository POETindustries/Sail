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
		http.Handle("/files/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/js/", http.FileServer(http.Dir(config.Cwd)))
		http.Handle("/theme/", http.FileServer(http.Dir(config.Cwd)))
		http.ListenAndServe(":8080", nil)
	}
}

func frontendHandler(writer http.ResponseWriter, req *http.Request) {
	t1 := time.Now().Nanosecond()

	if psqldb.Instance().Verify() {
		pages.Serve(pages.NewWithURL(req.URL.RequestURI())).WriteTo(writer)
	}

	t2 := time.Now().Nanosecond()
	fmt.Printf("Time to serve page: %d ms\n", (t2-t1)/1000000)
}

func backendHandler(writer http.ResponseWriter, req *http.Request) {
	// TODO check for session cookie, show login page if not present
}
