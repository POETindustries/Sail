package content

import "sail/page/meta"

// Content contains the information needed to generate a web page for display.
// This is the basic struct that contains all information needed to generate
// a correct and complete html page. It is the responsibility of the other
// functions and methods in package page to make sure its fields are
// properly initialized.
type Content struct {
	ID         uint32
	Title      string
	URL        string
	Content    string
	Meta       *meta.Meta
	TemplateID uint32

	Status int8
	Owner  string
	CDate  string
	EDate  string
}

// New creates a new Content object with usable defaults.
func New() *Content {
	return &Content{
		Meta:       meta.New(),
		TemplateID: 1}
}

func ByID(id uint32) *Content {
	cs := fromStorageByID(id)
	if len(cs) < 1 {
		if id == 1 {
			return NotFound()
		}
		return ByID(1)
	}
	return cs[0]
}

func ByIDs(ids ...uint32) []*Content {
	return nil
}

func ByURL(url string) *Content {
	cs := fromStorageByURL(url)
	if len(cs) < 1 {
		return nil
	}
	return cs[0]
}

func ByURLs(urls ...string) []*Content {
	return nil
}

func NotFound() *Content {
	p := New()
	p.ID, p.Title = 0, "Sorry about that"
	return p
}
