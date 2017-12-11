package ui

import (
	"github.com/nsf/termbox-go"
)

const (
	uiWidth       = 70
	topPadding    = 4
	widgetPadding = 1
	attrDefault   = termbox.ColorDefault
)

// A UI is the container for a UI which is centered on the screen.
type UI struct {
	Widgets   []Widget
	Selection int
}

// New constructs a new UI.
func New(widgets ...Widget) *UI {
	ui := &UI{Widgets: widgets}
	ui.MoveSelection(-1)

	return ui
}

// Listen begins listening for keyboard input and
// re-renders the UI when any is detected.
func (ui *UI) Listen() error {
	if err := ui.Render(); err != nil {
		return err
	}

	var (
		events = make(chan termbox.Event)
		errors = make(chan error)
	)

	go ui.handleInput(events, errors)

main:
	for {
		select {
		case e := <-events:
			switch e.Key {
			case termbox.KeyEsc:
				break main

			case termbox.KeyArrowDown:
				ui.MoveSelection(-1)

			case termbox.KeyArrowUp:
				ui.MoveSelection(1)

			default:
				if ui.Selection < len(ui.Widgets) {
					ui.Widgets[ui.Selection].HandleEvent(e, ui)
				}
			}

		case err := <-errors:
			return err

		default:
			if err := ui.Render(); err != nil {
				return err
			}
		}
	}

	return nil
}

// MoveSelection moves the selection up n amount of
// components. Use a negative value to move down.
func (ui *UI) MoveSelection(n int) {
	if ui.Selection < len(ui.Widgets) {
		ui.Widgets[ui.Selection].Deselect()
	}

	ui.Selection -= n

	if ui.Selection < 0 {
		ui.Selection = len(ui.Widgets) - 1
	}

	if ui.Selection >= len(ui.Widgets) {
		ui.Selection = 0
	}

	if !ui.Widgets[ui.Selection].IsSelectable() && ui.anySelectables() {
		step := 1
		if n < 0 {
			step = -1
		}

		ui.MoveSelection(step)
	}

	ui.Widgets[ui.Selection].Select()
}

func (ui *UI) anySelectables() bool {
	for _, widget := range ui.Widgets {
		if widget.IsSelectable() {
			return true
		}
	}

	return false
}

// Render renders all the widgets in order.
func (ui *UI) Render() error {
	termbox.Clear(attrDefault, attrDefault)

	var (
		w, _ = termbox.Size()
		x    = (w - uiWidth) / 2
		y    = topPadding
	)

	for i, widget := range ui.Widgets {
		y += widget.Render(x, y, w, i == ui.Selection) + widgetPadding
	}

	return termbox.Flush()
}

// handleInput starts an event loop.
func (ui *UI) handleInput(events chan termbox.Event, errors chan error) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			events <- ev

		case termbox.EventError:
			errors <- ev.Err
		}
	}
}

// A Widget is a widget in a UI, like a button or a text input.
type Widget interface {
	Render(x, y, w int, selected bool) int
	IsSelectable() bool
	HandleEvent(termbox.Event, *UI)
	Select()
	Deselect()
}
