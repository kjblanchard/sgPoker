package ui

import (
	"github.com/gdamore/tcell/v3"
	"github.com/kjblanchard/sgPoker/events"
)

type PlayAreaWindow struct {
	WindowBase
}

func (w *PlayAreaWindow) Initialize() {
	SetWindowSizeByPercent(&w.WindowBase, 0.80, 0.60, 2, 1)
	SetWindowLocationByPercent(&w.WindowBase, 0.10, 0.2, -1, 0)
}

func (w *PlayAreaWindow) Draw() {
	drawBox(w.Screen, w.X, w.Y, w.W+w.X, w.H+w.Y, tcell.StyleDefault, "Play")
}

func (w *PlayAreaWindow) HandleEvent(e *events.Event) {
}
