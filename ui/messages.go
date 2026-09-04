package ui

import (
	"github.com/gdamore/tcell/v3"
	"github.com/kjblanchard/sgPoker/events"
)

type MessageWindow struct {
	WindowBase
}

func (w *MessageWindow) Initialize() {
	SetWindowSizeByPercent(&w.WindowBase, 0.50, 0.20, 2, 1)
	SetWindowLocationByPercent(&w.WindowBase, 0.50, 0.0, -1, 0)
}

func (w *MessageWindow) Draw() {
	drawBox(w.Screen, w.X, w.Y, w.W+w.X, w.H+w.Y, tcell.StyleDefault, "Messages")
}

func (w *MessageWindow) HandleEvent(e *events.Event) {
}
