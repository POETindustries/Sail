package content

import "sail/object/meta"

// Content contains the information needed to generate a web page for display.
// This is the basic struct that contains all information needed to generate
// a correct and complete html page. It is the responsibility of the other
// functions and methods in package page to make sure its fields are
// properly initialized.
type Content struct {
	ID          uint32
	Title       string
	MachineName string
	Parent      uint32
	Status      int8
	Owner       string
	CDate       string
	EDate       string
	URL         string

	Content    string
	Meta       *meta.Meta
	TemplateID uint32
}

// New creates a new Content object with usable defaults.
func New() *Content {
	return &Content{
		Meta:       meta.New(),
		TemplateID: 1}
}

func ByID(id uint32) *Content {
	c := fromStorageMinByID(id)
	if c == nil {
		if id == 1 {
			return NotFound()
		}
		return ByID(1)
	}
	return c
}

func ByIDs(ids ...uint32) []*Content {
	return nil
}

func ByURL(url string) *Content {
	return fromStorageMinByURL(url)
}

func ByURLs(urls ...string) []*Content {
	return nil
}

func NotFound() *Content {
	p := New()
	p.ID, p.Title = 0, "Sorry about that"
	return p
}
