package pages

import (
	"io"
	"sail/domains"
	"sail/page"
	"sail/storage/pagestore"
)

// BuildWithURL expects a valid request uri in order to compile the
// corresponding page data.
//
// It is guaranteed to retun a functioning Page object even if the
// url parameter does not lead to any data.
func BuildWithURL(url string) *page.Page {
	var pages []*page.Page
	var err error
	if len(url) <= 1 {
		pages, err = fetchByID(1)
		if len(pages) == 0 || err != nil {
			return load404()
		}
	} else {
		pages, err = fetchByURL(url)
		if len(pages) == 0 || err != nil {
			pages, err = fetchByID(1)
			if len(pages) == 0 || err != nil {
				return load404()
			}
		}
	}
	pages[0].Domain = domains.BuildWithID(pages[0].Domain.ID)[0]
	return pages[0]
}

// Serve renders the page p and writes the result to the writer wr.
//
// If anything goes wrong, a non-nil error will be returned. In that
// case, it is the caller's responsibility to correct the contents of
// wr, which may have been partially written into.
func Serve(p *page.Page, wr io.Writer) error {
	return p.Domain.Template.Execute(wr, p)
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
