package pages

import (
	"bytes"
	"sail/cache"
	"sail/page"
	"sail/storage/pagestore"
	"sail/tmpl"
)

// Serve renders the page p and writes the result to the writer wr.
// If anything goes wrong, a non-nil error will be returned. In that
// case, it is the caller's responsibility to correct the contents of
// wr, which may have been partially written into.
func Serve(url string) *bytes.Buffer {
	if markup := cache.Instance().Markup(url); markup != nil {
		return bytes.NewBuffer(markup)
	}
	presenter := NewFromCache(url)
	if markup, err := presenter.Compile(); err == nil {
		cache.Instance().PushMarkup(url, markup.Bytes())
		return markup
	}
	return bytes.NewBufferString(tmpl.NOTFOUND404)
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
