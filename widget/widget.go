package widget

import (
	"fmt"
	"html/template"
)

// Widget holds all information to determine its type and function.
type Widget struct {
	ID      uint32
	Name    string
	RefName string
	Type    string
	Data    interface{}
}

// Menu returns the widget's data object cast to menu, if possible.
// It is guaranteed to return an object of the correct type; if the
// casting fails, an empty object is returned with all necessary
// components minimally initialized.
func (w *Widget) Menu() *Menu {
	m, ok := w.Data.(*Menu)
	if ok {
		return m
	}
	return &Menu{Entries: []*MenuEntry{}}
}

// Text returns the widget's data object cast to a text field, if
// possible. It is guaranteed to return an object of the correct type;
// if the casting fails, an empty object is returned with all
// necessary components minimally initialized.
func (w *Widget) Text() *Text {
	t, ok := w.Data.(*Text)
	if ok {
		return t
	}
	return &Text{Content: template.HTML("")}
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
