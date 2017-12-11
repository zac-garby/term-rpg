package ui

import (
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const inputWidth = 60

// An Input widget allows the user to enter a
// single line of text.
type Input struct {
	Text string
	Gap  int

	cursor int

	selectable
}

// Render renders the input to the screen at the
// given position. Returns the y offset of the
// next widget, relative to this one.
func (i *Input) Render(x, y, w int, selected bool) int {
	x += uiWidth/2 - inputWidth/2

	space := strings.Repeat(" ", inputWidth-runewidth.StringWidth(i.Text))
	print(x, y, -1, i.Text+space, termbox.AttrReverse, termbox.AttrReverse)

	if selected && i.cursor >= 0 && i.cursor-1 < len(i.Text) {
		termbox.SetCell(x+i.cursor, y, rune((i.Text + " ")[i.cursor]), attrDefault, attrDefault)
	}

	return i.Gap
}

// HandleEvent enters some text into the input.
func (i *Input) HandleEvent(e termbox.Event, ui *UI) {
	switch e.Key {
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		i.Backspace()

	case termbox.KeyDelete:
		i.Delete()

	case termbox.KeySpace:
		i.Insert(" ")

	case termbox.KeyTab:
		i.Insert("    ")

	case termbox.KeyArrowLeft:
		if i.cursor > 0 {
			i.cursor--
		}

	case termbox.KeyArrowRight:
		if i.cursor < len(i.Text) {
			i.cursor++
		}

	case termbox.KeyEnter:
		ui.MoveSelection(-1)

	default:
		i.Insert(string(e.Ch))
	}
}

// Select moves the cursor back to the beginning
// of the input.
func (i *Input) Select() {
	i.cursor = 0
}

// Insert inserts some text into the input field
// and moves the cursor accordingly.
func (i *Input) Insert(s string) {
	if len(i.Text+s) < inputWidth {
		i.Text = i.Text[:i.cursor] + s + i.Text[i.cursor:]
		i.cursor += len(s)
	}
}

// Backspace removes the character at the cursor.
func (i *Input) Backspace() {
	if len(i.Text) >= 0 && i.cursor > 0 {
		i.Text = i.Text[:i.cursor-1] + i.Text[i.cursor:]
		i.cursor--
	}
}

// Delete removes the character after the cursor.
// Backspace removes the character at the cursor.
func (i *Input) Delete() {
	if len(i.Text) >= 0 && i.cursor < len(i.Text) {
		i.Text = i.Text[:i.cursor] + i.Text[i.cursor+1:]
	}
}
