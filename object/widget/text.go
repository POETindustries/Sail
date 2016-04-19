package widget

// Text implements WidgetData. It holds arbitrary text for display.
type Text struct {
	Content string
}

func (t *Text) Copy() Data {
	return &Text{Content: t.Content}
}
