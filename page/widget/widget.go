package widget

import "fmt"

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
func New() *Widget {
	return &Widget{}
}

// String prints the widget's data in an easily readable format.
func (w *Widget) String() string {
	str := "WIDGET '%s': {ID:%d | RefName:%s | Type:%s | Data:%+v}"
	return fmt.Sprintf(str, w.Name, w.ID, w.RefName, w.Type, w.Data)
}

// ByIDs returns widgets that match the given parameter(s).
//
// It should be used for fetching one or more widgets for rendering
// and is guaranteed to contain at least one correctly set up widget
// at the first position of the returned slice.
func ByIDs(ids ...uint32) []*Widget {
	ws := fromStorageByID(ids...)
	if len(ws) < 1 {
		return []*Widget{New()}
	}
	for _, w := range ws {
		fetchData(w)
	}
	return ws
}
