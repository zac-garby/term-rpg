package game

import "github.com/Zac-Garby/term-rpg/ui"

// Display replaces the widgets in a UI with the Game's
// WorldRenderer instance.
func (g *Game) Display(u *ui.UI) {
	u.Widgets = []ui.Widget{
		g.Renderer,
		&ui.Text{Text: "\nConnected to server"},
	}

	u.Selection = 0
}
