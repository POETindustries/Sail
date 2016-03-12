package page

import (
	"fmt"
	"sail/conf"
	"sail/errors"
	"sail/page/data"
	"sail/storage/templatestore"
)

// BuildWithID returns templates that match the given id(s).
//
// It should be used to prepare one or more templates for rendering
// and is guaranteed to contain at least one correctly set up template
// at the first position of the returned slice.
func BuildWithID(ids ...uint32) []*data.Template {
	templates, err := fetchTemplateByID(ids...)
	if err != nil || len(templates) < 1 {
		errors.Log(err, conf.Instance().DevMode)
		return []*data.Template{data.NewTemplate()}
	}
	for _, t := range templates {
		widgetIDs, err := fetchWidgetIDs(t.ID)
		if err != nil {
			errors.Log(err, conf.Instance().DevMode)
			t.WidgetIDs = []uint32{}
		}
		t.WidgetIDs = widgetIDs
		widgets := WidgetsWithID(t.WidgetIDs...)
		for _, w := range widgets {
			t.Widgets[w.RefName] = w
		}
		t.Parse()
		Cache().PushTemplate(t)
		fmt.Printf("template added to cache: %d\n", t.ID)
	}
	return templates
}

func TemplateFromCache(id uint32) *data.Template {
	if t := Cache().Template(id); t != nil {
		fmt.Printf("found template in cache: %d\n", id)
		return t
	}
	return BuildWithID(id)[0]
}

func fetchTemplateByID(ids ...uint32) ([]*data.Template, error) {
	return templatestore.Get().ByID(ids...).Templates()
}

func fetchWidgetIDs(id uint32) ([]uint32, error) {
	return templatestore.Get().ByID(id).WidgetIDs()
}

func fetchWidgets(ids ...uint32) []*data.Widget {
	return WidgetsWithID(ids...)
}
