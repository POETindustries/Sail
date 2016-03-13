package page

import "bytes"

// Presenter initiates page creation and loading for handling requests
// by users from the www. It also serves as the content provider for
// templates.
//
// All exported functions and fields that return strings and
// string-derived types are safe for use inside a template. All exported
// functions and fields of type bool are safe for use as conditions
// inside templates.
type Presenter interface {
	Compile() *bytes.Buffer
	Message() string
	SetMessage(msg string)
	URL() string
	SetURL(url string)
}
