package renderer

import (
	"github.com/Zac-Garby/term-rpg/global"
	"github.com/Zac-Garby/term-rpg/ui"
	"github.com/Zac-Garby/term-rpg/world"
	"github.com/nsf/termbox-go"
)

// A WorldRenderer is a UI widget which renders
// a world and sends events to a channel.
type WorldRenderer struct {
	World        *world.World
	RoomX, RoomY int

	// Events are received on the HandleEvent method
	// and immediately sent down the Events channel.
	Events chan termbox.Event
}

// New constructs a new WorldRenderer to render a
// World.
func New(w *world.World, rx, ry int) *WorldRenderer {
	return &WorldRenderer{
		World:  w,
		Events: make(chan termbox.Event),
		RoomX:  rx,
		RoomY:  ry,
	}
}

// IsSelectable always returns true.
func (wr *WorldRenderer) IsSelectable() bool { return true }

// Select is called when the renderer is selected.
func (wr *WorldRenderer) Select() {}

// Deselect is called when the renderer is deselected.
func (wr *WorldRenderer) Deselect() {}

// Render renders the world and the entities
// on top of it.
func (wr *WorldRenderer) Render(x, y, w int, selected bool) int {
	box(x-1, y-1, global.RoomWidth+1, global.RoomHeight+1, termbox.AttrBold, termbox.AttrBold)

	room := wr.World.Rooms[wr.RoomY][wr.RoomX]

	for i := 0; i < global.RoomWidth; i++ {
		for j := 0; j < global.RoomHeight; j++ {
			var (
				tile = room.Tiles[j][i]
				attr = TileColour(tile)
			)

			termbox.SetCell(i+x, j+y, ' ', termbox.ColorDefault, attr)
		}
	}

	for _, ent := range wr.World.Entities {
		if ent.GetRoomX() == wr.RoomX && ent.GetRoomY() == wr.RoomY {
			var (
				char = ent.Character()
				tile = room.Tiles[ent.GetY()][ent.GetX()]
				bg   = TileColour(tile)
				fg   = termbox.AttrBold | termbox.ColorRed
			)

			termbox.SetCell(ent.GetX()+x, ent.GetY()+y, char, fg, bg)
		}
	}

	return global.RoomHeight + 2
}

// HandleEvent sends an event to the Events
// channel.
func (wr *WorldRenderer) HandleEvent(e termbox.Event, u *ui.UI) {
	wr.Events <- e
}
