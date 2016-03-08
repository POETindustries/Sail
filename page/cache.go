package page

import (
	"sail/page/data"
)

type cache struct {
	pages     map[string]*data.Page
	templates map[uint32]*data.Template
	markup    map[string][]byte
}

var instance *cache

func Cache() *cache {
	if instance == nil {
		instance = new()
	}
	return instance
}

func (c *cache) Template(id uint32) *data.Template {
	return c.templates[id]
}

func (c *cache) Page(url string) *data.Page {
	return c.pages[url]
}

func (c *cache) Markup(url string) []byte {
	return c.markup[url]
}

func (c *cache) PushTemplate(t *data.Template) {
	c.templates[t.ID] = t
}

func (c *cache) PushPage(p *data.Page) {
	c.pages[p.URL] = p
}

func (c *cache) PushMarkup(url string, m []byte) {
	c.markup[url] = m
}

// PopDomain removes the domain with the given id from the cache.
// If there are cached pages that depend on the domain, they, too,
// are deleted from their respective caches.
func (c *cache) PopTemplate(id uint32) {
	var urls []string
	for _, p := range c.pages {
		if p.Template.ID == id {
			urls = append(urls, p.URL)
		}
	}
	for _, url := range urls {
		c.PopPage(url)
		c.PopMarkup(url)
	}
	delete(c.templates, id)
}

// PopPage removes the given page from the cache. If there is
// cached markup belogning to the page, it is also deleted from
// its cache.
func (c *cache) PopPage(url string) {
	delete(c.pages, url)
	c.PopMarkup(url)
}

// PopMarkup deletes all markup from the cache that match the url
func (c *cache) PopMarkup(url string) {
	delete(c.markup, url)
}

func new() *cache {
	c := cache{}
	c.pages = make(map[string]*data.Page)
	c.templates = make(map[uint32]*data.Template)
	c.markup = make(map[string][]byte)

	return &c
}
