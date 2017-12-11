package renderer

import "github.com/nsf/termbox-go"

func box(x, y, w, h int, fg, bg termbox.Attribute) {
	for i := x; i <= x+w; i++ {
		termbox.SetCell(i, y, borderTop, fg, bg)
		termbox.SetCell(i, y+h, borderTop, fg, bg)
	}

	for i := y; i <= y+h; i++ {
		termbox.SetCell(x, i, borderSide, fg, bg)
		termbox.SetCell(x+w, i, borderSide, fg, bg)
	}

	termbox.SetCell(x, y, borderTopLeft, fg, bg)
	termbox.SetCell(x+w, y, borderTopRight, fg, bg)
	termbox.SetCell(x+w, y+h, borderBottomRight, fg, bg)
	termbox.SetCell(x, y+h, borderBottomLeft, fg, bg)
}
