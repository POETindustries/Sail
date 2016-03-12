package widget

import (
	"fmt"
	"sail/conf"
	"sail/errors"
)

// Widget is a small piece of software that can be embedded into a web page.
// It performs specific tasks or holds specific information or functionality.
//
// There are several types of widgets, each designed to fulfill more or less
// specific tasks.
//
// Menu
//
// On a usual webpage, menus are probably the most common widgets.
// They display a list of clickable elements, each directing the user
// to another location.
//
// Sail's menus currently don't support submenus.
//
// Text Field
//
// Text fields are generally used to convey a compact amount of
// information that should be displayed in more than one location
// and at the same time does not really fit well within the contents
// of any one page. Contact information in the sidebar is one typical
// use case for a text field widget.
type Widget struct {
	ID      uint32
	Name    string
	RefName string
	Type    string
	Data    interface{}
}

// NewWidget creates and returns a new widget object.
func NewWidget() *Widget {
	return &Widget{}
}

// String prints the widget's data in an easily readable format.
func (w *Widget) String() string {
	str := "WIDGET '%s': {ID:%d | RefName:%s | Type:%s | Data:%+v}"
	return fmt.Sprintf(str, w.Name, w.ID, w.RefName, w.Type, w.Data)
}

// Menu contains ordered, clickable elements.
type Menu struct {
	Entries []*MenuEntry
}

// String prints the menu's data in an easily readable format.
func (m *Menu) String() string {
	return fmt.Sprintf("MENU: {%+v}", m.Entries)
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

// Text implements WidgetData. It holds arbitrary text for display.
type Text struct {
	Content string
}

// WidgetWithID returns widgets that match the given parameter(s).
//
// It should be used for fetching one or more widgets for rendering
// and is guaranteed to contain at least one correctly set up widget
// at the first position of the returned slice.
func WidgetsWithID(id ...uint32) []*data.Widget {
	w, err := fetchWidgetByID(id...)
	if err != nil || len(w) < 1 {
		errors.Log(err, conf.Instance().DevMode)
		return []*data.Widget{data.NewWidget()}
	}
	if err = fetchWidgetData(w); err != nil {
		errors.Log(err, conf.Instance().DevMode)
	}
	return w
}

func fetchWidgetByID(id ...uint32) ([]*data.Widget, error) {
	return widgetstore.Get().ByID(id...).Widgets()
}

func fetchWidgetData(widgets []*data.Widget) (err error) {
	for _, w := range widgets {
		switch w.Type {
		case "menu":
			w.Data, err = fetchMenuData(w.ID)
		case "text":
			w.Data, err = fetchTextData(w.ID)
		}
		if err != nil {
			w.Data = nil
			return
		}
	}
	return
}

func fetchMenuData(id uint32) (*data.Menu, error) {
	return widgetstore.Get().ByID(id).Ascending().Menu()
}

func fetchTextData(id uint32) (*data.Text, error) {
	return widgetstore.Get().ByID(id).TextField()
}
