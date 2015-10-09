package widget

import (
	"fmt"
)

// Widget holds all information to determine its type and function.
type Widget struct {
	ID      uint32
	Name    string
	RefName string
	Type    string
	Data    interface{}
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
