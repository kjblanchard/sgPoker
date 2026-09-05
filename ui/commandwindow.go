package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v3"
	"github.com/kjblanchard/sgPoker/events"
)

type CommandWindow struct {
	WindowBase
	typingString string
	cursorX      int
	cursorY      int
}

func (w *CommandWindow) Initialize() {
	SetWindowSizeByPercent(&w.WindowBase, 0.80, 0.20, 2, 1)
	SetWindowLocationByPercent(&w.WindowBase, 0.00, 0.80, 0, 0)
	w.cursorX = w.X + 2
	w.cursorY = w.Y + 1
}

func (w *CommandWindow) Draw() {
	drawBox(w.Screen, w.X, w.Y, w.W+w.X, w.H+w.Y, tcell.StyleDefault, "Commands")
	drawText(w.Screen, w.X+1, w.Y+1, w.X+2, w.Y+1, tcell.StyleDefault, ">")
	drawText(w.Screen, w.X+2, w.Y+1, w.W-2, w.H-2, tcell.StyleDefault, w.typingString)
	// w.Screen.ShowCursor(w.cursorX+len(w.typingString)/w.W, w.cursorY%w.H)
}

func (w *CommandWindow) HandleEvent(e *events.Event) {
	//Handle tcell types
	if e.Type == events.EventTypeTCell {
		et, ok := e.Data.(tcell.Event)
		if !ok {
			return
		}
		// Handle if we should quit, resize event, or  something else.
		switch ev := et.(type) {
		case *tcell.EventKey:
			w.typingString += string(ev.Str())
			e := events.Event{Type: events.EventTypeStatusMessage, Data: fmt.Sprintf("Just pressed %s making %s", ev.Str(), w.typingString)}
			events.EventBus <- e
		}

	}
}
