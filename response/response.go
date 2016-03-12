package response

import (
	"bytes"
	"fmt"
	"net/http"
	"sail/page"
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

func (r *Response) FromCache() *Presenter {
	if p := page.Cache().Page(r.request.URL.RequestURI()); p != nil {
		fmt.Printf("page found in cache: %d\n", p.ID)
		return &Presenter{page: p, markup: bytes.NewBufferString("")}
	}
	return r.FromURL(true)
}

// NewWithID expects an id value in order to compile the
// corresponding page data. It is guaranteed to retun a functioning
// presenter object even if the id parameter does not lead to any data.
func (r *Response) FromID(id uint32) *Presenter {
	presenter := NewPresenter()
	pages, err := page.FetchPageByID(id)
	if len(pages) == 0 || err != nil {
		presenter.page = page.Load404()
	} else {
		pages[0].Template = page.TemplateFromCache(pages[0].Template.ID)
		presenter.page = pages[0]
		presenter.url = pages[0].URL
	}
	return presenter
}

// NewWithURL expects a valid request uri in order to compile the
// corresponding page data. It is guaranteed to retun a functioning
// presenter object even if the url parameter does not lead to any data.
func (r *Response) FromURL(cacheEnabled bool) *Presenter {
	presenter := NewPresenter()
	url := r.request.URL.RequestURI()
	if len(url) <= 1 {
		return r.FromID(presenter.FallbackID)
	}
	presenter.url = url
	pages, err := page.FetchPageByURL(url)
	if len(pages) == 0 || err != nil {
		return r.FromID(presenter.FallbackID)
	}
	pages[0].Template = page.TemplateFromCache(pages[0].Template.ID)
	presenter.page = pages[0]
	if cacheEnabled {
		page.Cache().PushPage(pages[0])
		fmt.Printf("page added to cache: %d\n", pages[0].ID)
	}
	return presenter
}

func (r *Response) Serve() {
	buf := bytes.NewBufferString(data.NOTFOUND404)
	if markup := page.Cache().Markup(r.URL); markup != nil {
		buf = bytes.NewBuffer(markup)
	}
	presenter := r.FromCache()
	if mk, err := presenter.Compile(); err == nil {
		page.Cache().PushMarkup(presenter.url, mk.Bytes())
		buf = mk
	}
	buf.WriteTo(r.response)
}
