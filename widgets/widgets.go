package widgets

import (
	"sail/conf"
	"sail/errors"
	"sail/storage/widgetstore"
	"sail/widget"
)

// BuildWithID returns widgets that match the given parameter(s).
//
// It should be used for fetching one or more widgets for rendering
// and is guaranteed to contain at least one correctly set up widget
// at the first position of the returned slice.
func BuildWithID(id ...uint32) []*widget.Widget {
	w, err := fetchByID(id...)
	if err != nil || len(w) < 1 {
		errors.Log(err, conf.Instance().DevMode)
		return []*widget.Widget{widget.New()}
	}
	if err = fetchData(w); err != nil {
		errors.Log(err, conf.Instance().DevMode)
	}
	return w
}

func fetchByID(id ...uint32) ([]*widget.Widget, error) {
	return widgetstore.Get().ByID(id...).Widgets()
}

func fetchData(widgets []*widget.Widget) (err error) {
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

func fetchMenuData(id uint32) (*widget.Menu, error) {
	return widgetstore.Get().ByID(id).Ascending().Menu()
}

func fetchTextData(id uint32) (*widget.Text, error) {
	return widgetstore.Get().ByID(id).TextField()
}
