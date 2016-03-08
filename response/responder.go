package response

import (
	"bytes"
	"fmt"
	"net/http"
	"sail/page"
	"sail/page/data"
)

// Responder handles all preparation and execution of respnses
// to http requests
type Responder struct {
	FallbackURL string
	FallbackID  uint32
	presenter   *Presenter
	request     *http.Request
}

func NewResponder(req *http.Request) *Responder {
	return &Responder{
		FallbackID:  1,
		FallbackURL: "/home",
		request:     req}
}

func (r *Responder) FromCache() *Presenter {
	if p := page.Cache().Page(r.request.URL.RequestURI()); p != nil {
		fmt.Printf("page found in cache: %d\n", p.ID)
		return &Presenter{page: p, markup: bytes.NewBufferString("")}
	}
	return r.FromURL(true)
}

// NewWithURL expects a valid request uri in order to compile the
// corresponding page data. It is guaranteed to retun a functioning
// presenter object even if the url parameter does not lead to any data.
func (r *Responder) FromURL(cacheEnabled bool) *Presenter {
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

// NewWithID expects an id value in order to compile the
// corresponding page data. It is guaranteed to retun a functioning
// presenter object even if the id parameter does not lead to any data.
func (r *Responder) FromID(id uint32) *Presenter {
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

// Serve renders the page p and writes the result to the writer wr.
// If anything goes wrong, a non-nil error will be returned. In that
// case, it is the caller's responsibility to correct the contents of
// wr, which may have been partially written into.
func (r *Responder) Serve() *bytes.Buffer {
	if mk := page.Cache().Markup(r.request.URL.RequestURI()); mk != nil {
		return bytes.NewBuffer(mk)
	}
	presenter := r.FromCache()
	if mk, err := presenter.Compile(); err == nil {
		page.Cache().PushMarkup(presenter.url, mk.Bytes())
		return mk
	}
	return bytes.NewBufferString(data.NOTFOUND404)
}
