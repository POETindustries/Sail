package page

import (
	"sail/page/data"
	"sail/storage/pagestore"
)

func FetchPageByURL(urls ...string) ([]*data.Page, error) {
	return pagestore.Get().ByURL(urls...).Pages()
}

func FetchPageByID(ids ...uint32) ([]*data.Page, error) {
	return pagestore.Get().ByID(ids...).Pages()
}

func Load404() *data.Page {
	p := data.NewPage()
	p.ID, p.Title = 0, "Sorry about that"
	return p
}
