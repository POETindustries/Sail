package widget

import "fmt"

// Nav contains ordered, clickable elements that point to places
// within the website.
type Nav struct {
	Entries []*NavEntry
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

// String prints the entry's data in an easily readable format.
func (e *NavEntry) String() string {
	str := "ENTRY '%s': {ID:%d | Image:%s | RefID:%d | RefURL:%s | Pos:%d | Active:%t}"
	return fmt.Sprintf(str, e.Name, e.ID, e.Image, e.RefID, e.RefURL, e.Pos, e.Active)
}
