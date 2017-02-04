package cache

import (
	"sail/object/content"
	"sail/object/template"
)

type cache struct {
	content   map[string]*content.Content
	templates map[uint32]*template.Template
	markup    map[string][]byte
	idToURL   map[string]string
	urlToID   map[string]string
}

var instance *cache

// DB returns the cache instance.
func DB() *cache {
	if instance == nil {
		instance = new()
	}
	return instance
}

func (c *cache) ObjectID(url string) string {
	return c.urlToID[url]
}

func (c *cache) ObjectURL(id string) string {
	return c.idToURL[id]
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

func (c *cache) PushURL(id string, url string) {
	c.idToURL[id] = url
	c.urlToID[url] = id
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

func (c *cache) PopURL(id string) {
	u := c.idToURL[id]
	delete(c.idToURL, id)
	delete(c.urlToID, u)
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
	c.content = make(map[string]*content.Content, 16)
	c.templates = make(map[uint32]*template.Template, 16)
	c.markup = make(map[string][]byte, 16)
	c.idToURL = make(map[string]string, 16)
	c.urlToID = make(map[string]string, 16)

	return &c
}
