package cache

import (
	"sail/domain"
	"sail/page"
)

type cache struct {
	pages   map[string]*page.Page
	domains map[uint32]*domain.Domain
	markup  map[string][]byte
}

var instance *cache

func Instance() *cache {
	if instance == nil {
		instance = new()
	}
	return instance
}

func (c *cache) Domain(id uint32) *domain.Domain {
	return c.domains[id]
}

func (c *cache) Page(url string) *page.Page {
	return c.pages[url]
}

func (c *cache) Markup(url string) []byte {
	return c.markup[url]
}

func (c *cache) PushDomain(d *domain.Domain) {
	c.domains[d.ID] = d
}

func (c *cache) PushPage(p *page.Page) {
	c.pages[p.URL] = p
}

func (c *cache) PushMarkup(url string, m []byte) {
	c.markup[url] = m
}

// PopDomain removes the domain with the given id from the cache.
// If there are cached pages that depend on the domain, they, too,
// are deleted from their respective caches.
func (c *cache) PopDomain(id uint32) {
	var urls []string
	for _, p := range c.pages {
		if p.Domain.ID == id {
			urls = append(urls, p.URL)
		}
	}
	for _, url := range urls {
		c.PopPage(url)
		c.PopMarkup(url)
	}
	delete(c.domains, id)
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
	c.pages = make(map[string]*page.Page)
	c.domains = make(map[uint32]*domain.Domain)
	c.markup = make(map[string][]byte)

	return &c
}
