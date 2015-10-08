package widget

import "html/template"

// Text implements WidgetData. It holds arbitrary text for display.
type Text struct {
	Content template.HTML
}
