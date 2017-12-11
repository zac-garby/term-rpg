package ui

// A Text widget renders some text to the screen.
type Text struct {
	Text string
	Gap  int

	unselectable
}

// Render renders the text to the screen at the
// given position. Returns the y offset of the
// next widget relative to this one.
func (t *Text) Render(x, y, w int, selected bool) int {
	lines := print(x, y, uiWidth, t.Text, attrDefault, attrDefault)

	return t.Gap + lines - 1
}
