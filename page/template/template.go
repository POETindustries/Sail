package template

import (
	"fmt"
	"html/template"
	"io"
	"sail/conf"
	"sail/errors"
)

var funcMap = template.FuncMap{
	"even": even}

func even(val int) bool {
	return val%2 == 0
}

// Template is the data structure that contains all data necessary
// to render the template files and all widgets contained within.
type Template struct {
	ID        uint32
	Name      string
	WidgetIDs []uint32
	template  *template.Template
	Widgets   map[string]*Widget
}

// NewTemplate creates a new Template object
func NewTemplate() *Template {
	return &Template{
		Name:    "404",
		Widgets: make(map[string]*Widget)}
}

// Execute applies a parsed template to the specified data object,
// writing the output to wr. If an error occurs during execution, it
// is the responsibility of the caller to handle partially written
// output.
func (t *Template) Execute(wr io.Writer, data interface{}) (err error) {
	if t.template == nil {
		err = errors.NilPointer()
	} else {
		err = t.template.ExecuteTemplate(wr, "frame.html", data)
	}
	errors.Log(err, conf.Instance().DevMode)
	return
}

// Parse parses the template files pointed at by the template.
func (t *Template) Parse() {
	if t.Name == "404" {
		t.template, _ = template.New("frame").Parse(NOTFOUND404)
	} else {
		dir := conf.Instance().TmplDir + t.Name
		tpl, err := template.New("").Funcs(funcMap).ParseGlob(dir + "/*.html")
		if err != nil {
			errors.Log(err, conf.Instance().DevMode)
			tpl, _ = template.New("frame").Parse(NOTFOUND404)
		}
		t.template = tpl
	}
}

// String prints the template's data in an easily readable format.
func (t *Template) String() string {
	str := "TEMPLATE '%s': {ID:%d | WidgetIDs:%+v | Template:%+v | Widgets:%+v}"
	return fmt.Sprintf(str, t.Name, t.ID, t.WidgetIDs, t.template, t.Widgets)
}

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
