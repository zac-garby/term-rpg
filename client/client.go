package client

import (
	"fmt"
	"log"
	"net"

	"github.com/Zac-Garby/term-rpg/entity"
	"github.com/Zac-Garby/term-rpg/game"
	"github.com/Zac-Garby/term-rpg/global"
	"github.com/Zac-Garby/term-rpg/message"
	"github.com/Zac-Garby/term-rpg/ui"
	"github.com/satori/go.uuid"
)

// A Client is a client connected to the server.
type Client struct {
	Name    string
	Address string
	*game.Game
	ID uuid.UUID

	u        *ui.UI
	infoText *ui.Text

	conn net.Conn

	// Player is the player entity the client has
	// control of.
	Player *entity.Player
}

// New constructs a new Client from an address and
// a name.
func New(addr, name string, u *ui.UI) *Client {
	return &Client{
		Name:    name,
		Address: addr,
		u:       u,
	}
}

// Listen starts the client listening to the server
// located at Address.
func (c *Client) Listen() {
	conn, err := net.Dial("tcp", c.Address)
	if err != nil {
		log.Println("client: could not connect:", err)

		c.u.Widgets = []ui.Widget{
			&ui.Text{Text: "Could not connect! Press *ESC* to exit", Gap: 1},
			&ui.Text{Text: err.Error()},
		}

		return
	}

	c.conn = conn
	if err := c.sendClientInfo(); err != nil {
		log.Println("client:", err)
	}

	for {
		msg, err := message.Receive(conn)

		// Assume the error is because of a manual
		// disconnect, but leave the game regardless.
		if err != nil {
			log.Println("client: leaving because of error:", err)
			c.Leave()
			break
		}

		go c.handleMessage(msg)
	}
}

// handleMessage is called for each received
// message from the server. It's called in a
// new goroutine each time, so it doesn't
// matter how long it runs for.
func (c *Client) handleMessage(msg message.Message) {
	switch m := msg.(type) {
	case *message.GameState:
		c.ID = m.ID

		for _, ent := range m.World.Entities {
			if player, ok := ent.(*entity.Player); ok && player.ID == m.ID {
				c.Player = player
				break
			}
		}

		if c.Player != nil {
			// Initialise the game from the server's information.
			c.Game = game.New(m.World, 0, 0)

			go c.handleInput()
		} else {
			log.Fatalln("client: player with id", m.ID, "not found")
		}

		c.infoText = &ui.Text{}
		c.setInfo()

		c.Display(c.u)
		c.u.Widgets = append(c.u.Widgets, c.infoText)

	case *message.NewPlayer:
		c.World.Entities = append(c.World.Entities, m.Player)

	case *message.Disconnect:
		if ok := c.World.DeletePlayer(m.ID); !ok {
			log.Println("client: couldn't find player to delete")
		}

	case *message.Movement:
		player := c.World.GetPlayer(m.ID)

		switch m.Direction {
		case 0:
			player.Up()

		case 1:
			player.Down()

		case 2:
			player.Left()

		case 3:
			player.Right()
		}
	}
}

// sendClientInfo sends some information to the
// server which is necessary to join the game.
func (c *Client) sendClientInfo() error {
	log.Println("client: sending client info")

	if text, ok := c.u.Widgets[0].(*ui.Text); ok {
		text.Text = "Sending client info..."
	}

	return message.Send(&message.ClientInfo{
		Name: c.Name,
	}, c.conn)
}

// Leave tells the server the client is leaving
// the game, then drops the connection.
func (c *Client) Leave() error {
	log.Println("client: leaving game")

	_ = message.Send(&message.Disconnect{}, c.conn)
	return c.conn.Close()
}

// setInfo sets the content of the info text
// to a useful value.
func (c *Client) setInfo() {
	if c.infoText == nil {
		return
	}

	c.infoText.Text = fmt.Sprintf(
		"X: %v, Y: %v",
		c.Player.X+c.Player.RoomX*global.RoomWidth,
		c.Player.Y+c.Player.RoomY*global.RoomHeight,
	)
}
