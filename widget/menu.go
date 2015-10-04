package widget

import (
	"bytes"
	"fmt"
)

// Menu implements WidgetData. It ordered, clickable elements.
type Menu struct {
	Entries []*MenuEntry
}

// Markup returns the markup string for this widget data type. Do not
// call this directly from template files!
func (m *Menu) Markup(htmlTagID string) string {
	mk := bytes.NewBufferString("<ul class='menu' id='" + htmlTagID + "'>")
	for _, e := range m.Entries {
		mk.WriteString("<li><a href='" + e.RefURL + "'>" + e.Name + "</a></li>")
	}
	mk.WriteString("</ul>")
	return mk.String()
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
	RefID   uint32
	RefURL  string
	Submenu uint32
	Pos     uint16
}

// String prints the entry's data in an easily readable format.
func (e *MenuEntry) String() string {
	str := "ENTRY '%s': {ID:%d | RefID:%d | RefURL:%s | Pos:%d}"
	return fmt.Sprintf(str, e.Name, e.ID, e.RefID, e.RefURL, e.Pos)
}
