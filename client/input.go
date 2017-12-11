package client

import (
	"log"

	"github.com/Zac-Garby/term-rpg/message"
	"github.com/nsf/termbox-go"
)

func (c *Client) handleInput() {
	for {
		e := <-c.Game.Renderer.Events

		switch e.Type {
		case termbox.EventKey:
			switch e.Ch {
			case 'w':
				go c.Up()

			case 's':
				go c.Down()

			case 'a':
				go c.Left()

			case 'd':
				go c.Right()
			}
		}
	}
}

// Up moves the player one position upwards.
func (c *Client) Up() {
	if c.Player == nil {
		return
	}

	if err := c.sendMove(0); err != nil {
		log.Println("client:", err)
		return
	}

	c.Player.Up()
	c.Renderer.RoomX = c.Player.RoomX
	c.Renderer.RoomY = c.Player.RoomY
	c.setInfo()
}

// Down moves the player one position downwards.
func (c *Client) Down() {
	if c.Player == nil {
		return
	}

	if err := c.sendMove(1); err != nil {
		log.Println("client:", err)
		return
	}

	c.Player.Down()
	c.Renderer.RoomX = c.Player.RoomX
	c.Renderer.RoomY = c.Player.RoomY
	c.setInfo()
}

// Left moves the player one position to the left.
func (c *Client) Left() {
	if c.Player == nil {
		return
	}

	if err := c.sendMove(2); err != nil {
		log.Println("client:", err)
		return
	}

	c.Player.Left()
	c.Renderer.RoomX = c.Player.RoomX
	c.Renderer.RoomY = c.Player.RoomY
	c.setInfo()
}

// Right moves the player one position to the right.
func (c *Client) Right() {
	if c.Player == nil {
		return
	}

	if err := c.sendMove(3); err != nil {
		log.Println("client:", err)
		return
	}

	c.Player.Right()
	c.Renderer.RoomX = c.Player.RoomX
	c.Renderer.RoomY = c.Player.RoomY
	c.setInfo()
}

func (c *Client) sendMove(dir uint8) error {
	return message.Send(&message.Movement{
		ID:        c.ID,
		Direction: dir,
	}, c.conn)
}
