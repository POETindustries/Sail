package page

import (
	"sail/conf"
	"sail/errors"
	"sail/page/data"
	"sail/storage/widgetstore"
)

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
