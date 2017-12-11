package ui

import (
	"unicode"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

// print prints the given text at (x, y) with the
// specified fg and bg attributes. It returns the
// amount of lines consumed (minimum of 1).
func print(x, y, wrap int, text string, fg, bg termbox.Attribute) int {
	var (
		startX    = x
		lines     = 1
		needsWrap = false
	)

	if wrap < 0 {
		wrap = 1 << 32
	}

	for _, c := range text {
		if c == '\\' {
			c = ' '
		}

		if c == '*' {
			fg ^= termbox.AttrBold
			continue
		}

		termbox.SetCell(x, y, c, fg, bg)

		x += runewidth.RuneWidth(c)
		if x > wrap+startX {
			needsWrap = true
		}

		if needsWrap && unicode.IsSpace(c) {
			x = startX
			y++
			lines++
			needsWrap = false
		}
	}

	return lines
}

type selectable struct{}

func (s selectable) IsSelectable() bool             { return true }
func (s selectable) HandleEvent(termbox.Event, *UI) {}
func (s selectable) Select()                        {}
func (s selectable) Deselect()                      {}

type unselectable struct{}

func (s unselectable) IsSelectable() bool             { return false }
func (s unselectable) HandleEvent(termbox.Event, *UI) {}
func (s unselectable) Select()                        {}
func (s unselectable) Deselect()                      {}
