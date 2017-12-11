package server

import (
	"github.com/Zac-Garby/term-rpg/message"
	"github.com/satori/go.uuid"
)

// GameState returns a GameState message to
// send to a client.
func (s *Server) GameState(id uuid.UUID) *message.GameState {
	return &message.GameState{
		ID:    id,
		World: s.World,
	}
}
