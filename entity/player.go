package entity

import (
	"encoding/gob"

	"github.com/Zac-Garby/term-rpg/global"
	"github.com/satori/go.uuid"
)

// A Player is an entity which is controlled by
// a real-life player.
type Player struct {
	Name string
	ID   uuid.UUID

	X, Y, RoomX, RoomY int
}

// NewPlayer constructs a new player with the
// specified parameters.
func NewPlayer(name string, id uuid.UUID, x, y, rx, ry int) *Player {
	return &Player{
		Name:  name,
		ID:    id,
		X:     x,
		Y:     y,
		RoomX: rx,
		RoomY: ry,
	}
}

// Character returns a smiley face character to
// render the player.
func (p *Player) Character() rune {
	return '▲'
}

// Up moves an entity one position upwards (negative Y).
func (p *Player) Up() {
	if p.Y > 0 {
		p.Y--
	} else if p.RoomY > 0 {
		p.RoomY--
		p.Y = global.RoomHeight - 1
	}
}

// Down moves an entity one position downwards (positive Y).
func (p *Player) Down() {
	if p.Y+1 < global.RoomHeight {
		p.Y++
	} else if p.RoomX < global.WorldHeight {
		p.RoomY++
		p.Y = 0
	}
}

// Left moves an entity one position to the left (negative X).
func (p *Player) Left() {
	if p.X > 0 {
		p.X--
	} else if p.RoomX > 0 {
		p.RoomX--
		p.X = global.RoomWidth - 1
	}
}

// Right moves an entity one position to the right (positive X).
func (p *Player) Right() {
	if p.X+1 < global.RoomWidth {
		p.X++
	} else if p.RoomX < global.WorldWidth {
		p.RoomX++
		p.X = 0
	}
}

// GetX gets the x-ordinatp.
func (p *Player) GetX() int { return p.X }

// GetY gets the y-ordinatp.
func (p *Player) GetY() int { return p.Y }

// GetRoomX gets the room x-ordinatp.
func (p *Player) GetRoomX() int { return p.RoomX }

// GetRoomY gets the room y-ordinatp.
func (p *Player) GetRoomY() int { return p.RoomY }

func init() {
	gob.Register(&Player{})
}
