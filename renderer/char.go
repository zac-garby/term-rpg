package renderer

import (
	"github.com/Zac-Garby/term-rpg/world"
	termbox "github.com/nsf/termbox-go"
)

const (
	borderSide        = '┃'
	borderTop         = '━'
	borderTopLeft     = '┏'
	borderTopRight    = '┓'
	borderBottomLeft  = '┗'
	borderBottomRight = '┛'
)

var tileColours = map[world.Tile]termbox.Attribute{
	world.Air:  termbox.ColorBlack,
	world.Wall: termbox.ColorWhite,
}

// TileColour gets the character to represent
// the given tile type.
func TileColour(t world.Tile) termbox.Attribute {
	ch, ok := tileColours[t]
	if ok {
		return ch
	}

	return '?'
}
