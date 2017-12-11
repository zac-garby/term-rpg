package server

import (
	"errors"
	"log"
	"net"
	"strings"

	"github.com/Zac-Garby/term-rpg/entity"
	"github.com/Zac-Garby/term-rpg/message"
	"github.com/Zac-Garby/term-rpg/world"
	"github.com/satori/go.uuid"
)

var (
	// ErrNoClient is returned when a client does not exist
	ErrNoClient = errors.New("server: the specified client does not exist")
)

type receivedMessage struct {
	msg    message.Message
	sender uuid.UUID
}

type connection struct {
	net.Conn
	name string
}

// A Server manages connections to multiple clients
// and stores the world they all exist in.
type Server struct {
	address     string
	connections map[uuid.UUID]*connection
	messages    chan receivedMessage

	World *world.World
}

// New constructs a new server on the given address.
func New(addr string) *Server {
	return &Server{
		address:     addr,
		connections: make(map[uuid.UUID]*connection),
		World:       world.New(),
	}
}

// Listen starts listening on the server's address
// for any incoming connections.
func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", ":"+s.address)
	if err != nil {
		return err
	}

	log.Println("server: listening on port", s.address)

	go s.handleMessages()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	id := uuid.NewV4()
	s.connections[id] = &connection{Conn: conn}

	go s.listenForMessages(id, s.connections[id])
}

func (s *Server) listenForMessages(id uuid.UUID, conn *connection) {
	msg, err := message.Receive(conn)
	if err != nil {
		log.Println("server:", err)
		return
	}

	if info, ok := msg.(*message.ClientInfo); ok {
		log.Println("server: user has connected:", info.Name)
		conn.name = info.Name

		player := entity.NewPlayer(
			info.Name,
			id,
			10, 5, // Position
			0, 0, // Room
		)

		s.World.Entities = append(s.World.Entities, player)

		if err := s.Send(id, s.GameState(id)); err != nil {
			if strings.Contains(err.Error(), "write: broken pipe") {
				s.handleDisconnect(id)
				return
			}

			log.Println("server:", err)
		}

		// Sent a NewPlayer message to all clients except
		// the newly joined one.
		if err := s.Broadcast(&message.NewPlayer{Player: player}, id); err != nil {
			log.Println("server:", err)
		}
	} else {
		log.Println("server: a client connected, but didn't send any client info")
		return
	}

	for {
		msg, err := message.Receive(conn)

		// An error was most likely caused by a disconnect,
		// but either way, end the connection.
		if err != nil {
			s.handleDisconnect(id)
			break
		}

		s.messages <- receivedMessage{
			msg:    msg,
			sender: id,
		}
	}
}

// Send sends a message to the client with the
// specified id.
func (s *Server) Send(id uuid.UUID, msg message.Message) error {
	conn, ok := s.connections[id]
	if !ok {
		return ErrNoClient
	}

	return message.Send(msg, conn)
}

// Broadcast broadcasts a message to all connected
// clients except the ones mentioned in the except
// parameter.
func (s *Server) Broadcast(msg message.Message, except ...uuid.UUID) error {
conns:
	for id := range s.connections {
		for _, exception := range except {
			if exception == id {
				continue conns
			}
		}

		if err := s.Send(id, msg); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) handleMessage(id uuid.UUID, msg message.Message) {
	switch msg := msg.(type) {
	case *message.Movement:
		player := s.World.GetPlayer(id)

		switch msg.Direction {
		case 0:
			player.Up()

		case 1:
			player.Down()

		case 2:
			player.Left()

		case 3:
			player.Right()
		}

		s.Broadcast(msg, id)
	}
}

func (s *Server) handleDisconnect(id uuid.UUID) {
	log.Println("server: disconnecting player:", id)

	s.Broadcast(&message.Disconnect{ID: id}, id)

	if ok := s.World.DeletePlayer(id); !ok {
		log.Println("server: couldn't find player to delete")
	}

	if err := s.connections[id].Close(); err != nil {
		log.Println("server: error in disconnect:", err)
	}
}

func (s *Server) handleMessages() {
	for {
		msg := <-s.messages
		go s.handleMessage(msg.sender, msg.msg)
	}
}
