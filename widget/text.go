package widget

// Text implements WidgetData. It holds arbitrary text for display.
type Text struct {
	Content string
}

// Markup returns the markup string for this widget data type. Do not
// call this directly from template files!
func (t *Text) Markup(htmlTagID string) string {
	return t.Content
}
