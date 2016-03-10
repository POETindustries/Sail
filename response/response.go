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

	presenter *Presenter
	response  http.ResponseWriter
	request   *http.Request
}

func New(wr http.ResponseWriter, req *http.Request) *Response {
	return &Response{
		FallbackID:  1,
		FallbackURL: "/home",
		URL:         req.URL.Path,
		response:    wr,
		request:     req}
}

func (r *Response) Serve() {

	if mk := page.Cache().Markup(r.URL)); mk != nil {
		return bytes.NewBuffer(mk)
	}
	presenter := r.FromCache()
	if mk, err := presenter.Compile(); err == nil {
		page.Cache().PushMarkup(presenter.url, mk.Bytes())
		return mk
	}
	return bytes.NewBufferString(data.NOTFOUND404)

	var buf *bytes.Buffer
	buf = bytes.NewBufferString(data.NOTFOUND404)
	buf.WriteTo(r.response)
}
