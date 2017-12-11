package world

import (
	"github.com/Zac-Garby/term-rpg/entity"
	"github.com/Zac-Garby/term-rpg/global"
	"github.com/satori/go.uuid"
)

// A Room is a single room in a world. Stores
// the tile data in a [roomHeight * roomWidth]
// matrix.
type Room struct {
	Tiles [global.RoomHeight][global.RoomWidth]uint8
}

// A World contains a matrix of rooms. There
// are [worldHeight * worldWidth] rooms.
type World struct {
	Rooms    [global.WorldHeight][global.WorldWidth]Room
	Entities []entity.Entity
}

// New constructs a new, empty, world.
func New() *World {
	return &World{}
}

// GetPlayer returns the player, if any, which
// has the id.
func (w *World) GetPlayer(id uuid.UUID) *entity.Player {
	for _, e := range w.Entities {
		if player, ok := e.(*entity.Player); ok && player.ID == id {
			return player
		}
	}

	return nil
}

// DeletePlayer deletes the player specified by
// id if it exists, and returns false if it doesn't.
func (w *World) DeletePlayer(id uuid.UUID) bool {
	index := -1

	for i, e := range w.Entities {
		if player, ok := e.(*entity.Player); ok && player.ID == id {
			index = i
			break
		}
	}

	if index < 0 {
		return false
	}

	copy(w.Entities[index:], w.Entities[index+1:])
	w.Entities[len(w.Entities)-1] = nil
	w.Entities = w.Entities[:len(w.Entities)-1]

	return true
}
