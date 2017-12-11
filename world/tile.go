package world

// A Tile is an 8-bit integer which represents
// a tile in a room.
type Tile = uint8

const (
	// Air tiles are invisible and passable.
	Air Tile = iota

	// Wall tiles are used for walls of buildings.
	Wall
)
