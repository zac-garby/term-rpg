package server

import (
	"fmt"

	"github.com/Zac-Garby/term-rpg/ui"
)

// Display overrides the ui's widgets to
// show a couple of messages.
func (s *Server) Display(u *ui.UI) {
	u.Widgets = []ui.Widget{
		&ui.Text{Text: fmt.Sprintf("Listening on *%s*", s.address)},
		&ui.Text{Text: "Press *ESC* to exit"},
	}
}
