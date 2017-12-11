package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Zac-Garby/term-rpg/client"
	"github.com/Zac-Garby/term-rpg/server"
	"github.com/Zac-Garby/term-rpg/ui"
	"github.com/nsf/termbox-go"
)

var (
	u *ui.UI
	f *os.File
)

func main() {
	var err error

	f, err = os.OpenFile("log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	log.SetOutput(f)
	log.Println("*** started at", time.Now())

	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}

	defer termbox.Close()

	u = &ui.UI{}
	toMainMenu()

	if err := u.Listen(); err != nil {
		log.Fatal(err)
	}
}

func exit() {
	termbox.Close()
	f.Close()
	os.Exit(0)
}

func toMainMenu() {
	u.Widgets = []ui.Widget{
		&ui.Text{Text: "*Welcome to term-rpg* by Zac Garby <me@zacgarby.co.uk>", Gap: 1},
		&ui.Text{Text: `Press *ESC* at any time to exit the game, use\
						*UP* and *DOWN* to navigate the menu, and\
						press *RETURN* to select an option.`, Gap: 1},

		&ui.Button{Text: "*join a server*", Callback: toJoin},
		&ui.Button{Text: "*host a new server*", Callback: toHost},
		&ui.Button{Text: "*change settings*", Callback: func() {}},
		&ui.Button{Text: "*exit the game*", Gap: 1, Callback: exit},
	}

	u.Selection = 0
	u.MoveSelection(-1)
}

func toJoin() {
	var (
		nameInput = &ui.Input{Gap: 2}
		addrInput = &ui.Input{Text: "localhost:51515", Gap: 2}
	)

	u.Widgets = []ui.Widget{
		&ui.Text{Text: "What's your name?", Gap: 1},
		nameInput,

		&ui.Text{Text: "What server would you like to connect to?", Gap: 1},
		addrInput,

		&ui.Button{Text: "join", Callback: func() { joinServer(nameInput.Text, addrInput.Text) }},
		&ui.Button{Text: "back", Callback: toMainMenu},
	}

	u.Selection = 0
	u.MoveSelection(-1)
}

func toHost() {
	addrInput := &ui.Input{Text: "51515", Gap: 2}

	u.Widgets = []ui.Widget{
		&ui.Text{Text: "What port would you like to host on?", Gap: 1},
		addrInput,

		&ui.Button{Text: "host", Callback: func() { startServer(addrInput.Text) }},
		&ui.Button{Text: "back", Callback: toMainMenu},
	}

	u.Selection = 0
	u.MoveSelection(-1)
}

func startServer(addr string) {
	s := server.New(addr)

	go s.Listen()
	s.Display(u)
}

func joinServer(name, addr string) {
	c := client.New(addr, name, u)

	u.Widgets = []ui.Widget{
		&ui.Text{Text: fmt.Sprintf("*Attempting to connect to %s*", addr), Gap: 1},
		&ui.Text{Text: "Press *ESC* to give up"},
	}

	go c.Listen()
}
