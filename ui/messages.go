package ui

import (
	"slices"

	"github.com/gdamore/tcell/v3"
	"github.com/kjblanchard/sgPoker/events"
)

type MessageWindow struct {
	WindowBase
	messages []string
}

func (w *MessageWindow) Initialize() {
	SetWindowSizeByPercent(&w.WindowBase, 0.50, 0.20, 2, 1)
	SetWindowLocationByPercent(&w.WindowBase, 0.50, 0.0, -1, 0)
}

func (w *MessageWindow) Draw() {
	drawBox(w.Screen, w.X, w.Y, w.W+w.X, w.H+w.Y, tcell.StyleDefault, "Messages")
	var i int = 0
	for _, m := range slices.Backward(w.messages) {
		drawText(w.Screen, w.X+1, w.Y+1+i, w.X+w.W-2, w.H-1, tcell.StyleDefault, m)
		i++
	}
}

func (w *MessageWindow) HandleEvent(e *events.Event) {
	if e.Type == events.EventTypeStatusMessage {
		s, ok := e.Data.(string)
		if !ok {
			return
		}
		//TODO this backing array probably grows forever, probably make this better one day. for now, make the slice a queue by removing first element
		if len(w.messages) > w.H-2 {
			w.messages = w.messages[1:]
		}
		w.messages = append(w.messages, s)
	}
}
