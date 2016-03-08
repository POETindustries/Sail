package response

import (
	"bytes"
	"net/http"
	"sail/page/data"
)

type Response struct {
	FallbackID  uint32
	FallbackURL string
	Message     string
	URL         string

	writer  http.ResponseWriter
	request *http.Request
}

func New(wr http.ResponseWriter, req *http.Request) *Response {
	return &Response{
		FallbackID:  1,
		FallbackURL: "/home",
		URL:         req.URL.Path,
		writer:      wr,
		request:     req}
}

func (r *Response) Serve() {
	var buf *bytes.Buffer
	buf = bytes.NewBufferString(data.NOTFOUND404)
	buf.WriteTo(r.writer)
}
