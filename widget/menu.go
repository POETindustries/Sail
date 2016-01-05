package widget

import "fmt"

// Menu contains ordered, clickable elements.
type Menu struct {
	Entries []*MenuEntry
}

// String prints the menu's data in an easily readable format.
func (m *Menu) String() string {
	str := "MENU: {%+v}"
	return fmt.Sprintf(str, m.Entries)
}

// MenuEntry contains all information about a specific menu entry.
type MenuEntry struct {
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
func (e *MenuEntry) String() string {
	str := "ENTRY '%s': {ID:%d | Image:%s | RefID:%d | RefURL:%s | Pos:%d | Active:%t}"
	return fmt.Sprintf(str, e.Name, e.ID, e.Image, e.RefID, e.RefURL, e.Pos, e.Active)
}
