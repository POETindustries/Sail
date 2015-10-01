package pages

import (
	"sail/conf"
	"sail/errors"
	"sail/page"
	"sail/storage/pagestore"
)

// BuildWithURL expects a valid request uri in order to compile the
// corresponding page data.
//
// It is guaranteed to retun a functioning Page object even if the
// passed string does not lead to any data.
func BuildWithURL(url string) *page.Page {
	p := page.New()
	if !fetchByURL(url, p) {
		return load404()
	}
	return p
}

func fetchByURL(url string, p *page.Page) bool {
	if len(url) <= 1 || pagestore.Get().ByURL(url).Execute(p) != nil {
		return fetchByID(1, p)
	}
	return true
}

func fetchByID(id uint32, p *page.Page) bool {
	if err := pagestore.Get().ByID(id).Execute(p); err != nil {
		errors.Log(err, conf.Instance().DevMode)
		return false
	}
	return true
}

func load404() *page.Page {
	p := page.New()
	p.ID, p.Title = 0, "Sorry about that"
	return p
}
