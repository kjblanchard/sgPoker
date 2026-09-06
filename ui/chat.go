package ui

import (
	"github.com/kjblanchard/sgPoker/events"
)

type ChatWindow struct {
	WindowBase
}

func (w *ChatWindow) Initialize() {
	SetWindowSizeByPercent(&w.WindowBase, 0.50, 0.20, 2, 1)
	SetWindowLocationByPercent(&w.WindowBase, 0.00, 0.0, -1, 0)
}

func (w *ChatWindow) Draw() {
	w.drawBoxOutline("Chat")
}

func (w *ChatWindow) HandleEvent(e *events.Event) {
}
