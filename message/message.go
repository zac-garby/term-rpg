package message

import (
	"encoding/gob"

	"github.com/Zac-Garby/term-rpg/entity"
	"github.com/Zac-Garby/term-rpg/world"
	"github.com/satori/go.uuid"
)

// A Message is some data which is sent down a
// TCP connection to tell the server or client
// something.
type Message interface{}

// GameState is sent from the server to the
// client. Contains information about the
// state of the server and gives the client
// it's UUID.
type GameState struct {
	ID    uuid.UUID
	World *world.World
}

// NewPlayer is sent from the server to all
// clients except the new player when a player
// connects to the server. This allows the
// clients to add a new entity for the player.
type NewPlayer struct {
	Player *entity.Player
}

// ClientInfo is sent from the client to the
// server. It contains information about the
// client.
type ClientInfo struct {
	Name string
}

// Disconnect is sent from the client to the
// server to tell it that he's disconnecting.
type Disconnect struct {
	ID uuid.UUID
}

// Movement is sent from the client to the
// server to tell it that he's moved in a
// direction.
//
// 0 -> up
// 1 -> down
// 2 -> left
// 3 -> right
//
// If sent from the server, ID is the id of
// the player who moved.
type Movement struct {
	ID        uuid.UUID
	Direction uint8
}

// init registers all the possible message
// types so they can be decoded.
func init() {
	gob.Register(&GameState{})
	gob.Register(&NewPlayer{})
	gob.Register(&ClientInfo{})
	gob.Register(&Disconnect{})
	gob.Register(&Movement{})
}
