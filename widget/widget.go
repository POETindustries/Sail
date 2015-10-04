package widget

import (
	"fmt"
	"html/template"
)

// Data is the interface that all types must implement if they
// are to be used for holding Widget's data.
type Data interface {
	Markup(htmlTagID string) string
}

// Widget holds all information to determine its type and function.
type Widget struct {
	ID      uint32
	Name    string
	RefName string
	Type    string
	Data    Data
}

// Markup returns the widget's data in a form fit for display inside
// an html document.
//
// DO NOT EVER try to call the widget data's markup() method directly!
// That method's return string is not suited for display inside an
// html document and Bad Things Will Happen®. Always use this, the
// widget's own Markup() method when the result is used for embedding
// in html code.
func (w *Widget) Markup() template.HTML {
	if w.Data == nil {
		return template.HTML("")
	}
	return template.HTML(w.Data.Markup(w.RefName))
}

// String prints the widget's data in an easily readable format.
func (w *Widget) String() string {
	str := "WIDGET '%s': {ID:%d | RefName:%s | Type:%s | Data:%+v}"
	return fmt.Sprintf(str, w.Name, w.ID, w.RefName, w.Type, w.Data)
}

// New creates and returns a new widget object.
func New() *Widget {
	return &Widget{}
}
