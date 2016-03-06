package templates

import (
	"fmt"
	"sail/cache"
	"sail/conf"
	"sail/errors"
	"sail/storage/templatestore"
	"sail/tmpl"
	"sail/widget"
	"sail/widgets"
)

// BuildWithID returns templates that match the given id(s).
//
// It should be used to prepare one or more templates for rendering
// and is guaranteed to contain at least one correctly set up template
// at the first position of the returned slice.
func BuildWithID(ids ...uint32) []*tmpl.Template {
	templates, err := fetchByID(ids...)
	if err != nil || len(templates) < 1 {
		errors.Log(err, conf.Instance().DevMode)
		return []*tmpl.Template{tmpl.New()}
	}
	for _, t := range templates {
		widgetIDs, err := fetchWidgetIDs(t.ID)
		if err != nil {
			errors.Log(err, conf.Instance().DevMode)
			t.WidgetIDs = []uint32{}
		}
		t.WidgetIDs = widgetIDs
		widgets := widgets.BuildWithID(t.WidgetIDs...)
		for _, w := range widgets {
			t.Widgets[w.RefName] = w
		}
		t.Compile()
		cache.Instance().PushTemplate(t)
		fmt.Printf("template added to cache: %d\n", t.ID)
	}
	return templates
}

func FromCache(id uint32) *tmpl.Template {
	if t := cache.Instance().Template(id); t != nil {
		fmt.Printf("found template in cache: %d\n", id)
		return t
	}
	return BuildWithID(id)[0]
}

func fetchByID(ids ...uint32) ([]*tmpl.Template, error) {
	return templatestore.Get().ByID(ids...).Templates()
}

func fetchWidgetIDs(id uint32) ([]uint32, error) {
	return templatestore.Get().ByID(id).WidgetIDs()
}

func fetchWidgets(ids ...uint32) []*widget.Widget {
	return widgets.BuildWithID(ids...)
}
