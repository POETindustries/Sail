package widget

import "fmt"

// Nav contains ordered, clickable elements that point to places
// within the website.
type Nav struct {
	Entries []*NavEntry
}

func (n *Nav) Copy() Data {
	var es []*NavEntry
	for _, e := range n.Entries {
		es = append(es, e.Copy())
	}
	return &Nav{Entries: es}
}

// String prints the menu's data in an easily readable format.
func (m *Nav) String() string {
	return fmt.Sprintf("NAV: {%+v}", m.Entries)
}

// NavEntry contains all information about a specific menu entry.
type NavEntry struct {
	ID      uint32
	Name    string
	Image   string
	RefID   uint32
	RefURL  string
	Submenu uint32
	Pos     uint16
	Active  bool
}

func (e *NavEntry) Copy() *NavEntry {
	return &NavEntry{
		ID:      e.ID,
		Name:    e.Name,
		Image:   e.Image,
		RefID:   e.RefID,
		RefURL:  e.RefURL,
		Submenu: e.Submenu,
		Pos:     e.Pos,
		Active:  e.Active}
}

// String prints the entry's data in an easily readable format.
func (e *NavEntry) String() string {
	str := "ENTRY '%s': {ID:%d | Image:%s | RefID:%d | RefURL:%s | Pos:%d | Active:%t}"
	return fmt.Sprintf(str, e.Name, e.ID, e.Image, e.RefID, e.RefURL, e.Pos, e.Active)
}
