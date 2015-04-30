package tmpl

type Template struct {

	// Templates contains the names of all templates used by this one.
	// It includes its own name, alway at index 0.
	Templates []string

	// Files contains the complete file names of each template in Templates.
	// This field can be passed to Go's template.ParseFiles method.
	Files []string

	// Content contains the whole markup of the template.
	// It is not used for frontend functionality and is instead intended for
	// use when editing templates from the backend.
	Content string

	// Name is the template's unique identifier.
	// It is used in the database and for reference in other templates.
	Name string
}
