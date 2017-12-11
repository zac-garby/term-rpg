package game

import (
	"github.com/Zac-Garby/term-rpg/renderer"
	"github.com/Zac-Garby/term-rpg/world"
)

// The Game coordinates the client, the world, and
// the rendering of them both.
type Game struct {
	*world.World
	Renderer *renderer.WorldRenderer
}

// New constructs a new game and connects the client
// to the given address.
func New(w *world.World, rx, ry int) *Game {
	return &Game{
		World:    w,
		Renderer: renderer.New(w, rx, ry),
	}
}
