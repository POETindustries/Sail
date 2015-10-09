package pages

import (
	"bytes"
	"sail/page"
	"sail/storage/pagestore"
	"sail/tmpl"
)

// Serve renders the page p and writes the result to the writer wr.
// If anything goes wrong, a non-nil error will be returned. In that
// case, it is the caller's responsibility to correct the contents of
// wr, which may have been partially written into.
func Serve(p *Presenter) *bytes.Buffer {
	if p.compile() != nil {
		markup := tmpl.NOTFOUND404
		return bytes.NewBufferString(markup)
	}
	return p.markup
}

func fetchByURL(urls ...string) ([]*page.Page, error) {
	return pagestore.Get().ByURL(urls...).Pages()
}

func fetchByID(ids ...uint32) ([]*page.Page, error) {
	return pagestore.Get().ByID(ids...).Pages()
}

func load404() *page.Page {
	p := page.New()
	p.ID, p.Title = 0, "Sorry about that"
	return p
}
