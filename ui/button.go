package ui

import (
	"github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

// A Button widget renders some text and calls
// a callback function when ENTER is pressed on
// it.
type Button struct {
	Text     string
	Callback func()
	Gap      int

	selectable
}

// Render renders the button to the screen at the
// given position. Returns the y offset of the
// next widget relative to this one.
func (b *Button) Render(x, y, w int, selected bool) int {
	var lines int
	x += uiWidth/2 - runewidth.StringWidth(b.Text)/2

	if selected {
		lines = print(x-2, y, uiWidth, "> "+b.Text, attrDefault, attrDefault)
	} else {
		lines = print(x, y, uiWidth, b.Text, attrDefault, attrDefault)
	}

	return b.Gap + lines - 1
}

// HandleEvent will call the button's callback
// function if the event is a keypress of the
// return key.
func (b *Button) HandleEvent(e termbox.Event, ui *UI) {
	if e.Key == termbox.KeyEnter {
		b.Callback()
	}
}
