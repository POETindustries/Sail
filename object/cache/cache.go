package cache

import (
	"sail/object/content"
	"sail/object/template"
)

type cache struct {
	content   map[string]*content.Content
	templates map[uint32]*template.Template
	markup    map[string][]byte
}

var instance *cache

func DB() *cache {
	if instance == nil {
		instance = new()
	}
	return instance
}

func (c *cache) Template(id uint32) *template.Template {
	return c.templates[id]
}

func (c *cache) Content(url string) *content.Content {
	return c.content[url]
}

func (c *cache) Markup(url string) []byte {
	return c.markup[url]
}

func (c *cache) PushTemplate(t *template.Template) {
	c.templates[t.ID] = t
}

func (c *cache) PushContent(p *content.Content) {
	c.content[p.URL] = p
}

func (c *cache) PushMarkup(url string, m []byte) {
	c.markup[url] = m
}

// PopTemplate removes the domain with the given id from the cache.
// If there is cached content that depends on the domain, they, too,
// are deleted from their respective caches.
func (c *cache) PopTemplate(id uint32) {
	var urls []string
	for _, p := range c.content {
		if p.TemplateID == id {
			urls = append(urls, p.URL)
		}
	}
	for _, url := range urls {
		c.PopContent(url)
		c.PopMarkup(url)
	}
	delete(c.templates, id)
}

// PopPage removes the given page from the cache. If there is
// cached markup belogning to the page, it is also deleted from
// its cache.
func (c *cache) PopContent(url string) {
	delete(c.content, url)
	c.PopMarkup(url)
}

// PopMarkup deletes all markup from the cache that match the url
func (c *cache) PopMarkup(url string) {
	delete(c.markup, url)
}

func new() *cache {
	c := cache{}
	c.content = make(map[string]*content.Content)
	c.templates = make(map[uint32]*template.Template)
	c.markup = make(map[string][]byte)

	return &c
}
