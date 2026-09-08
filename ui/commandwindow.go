package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v3"
	"github.com/kjblanchard/sgPoker/cards"
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
	w.drawBoxOutline("Commands")
	w.drawText(0, 0, 2, 2, ">")
	w.drawText(1, 0, 0, 0, w.typingString)
}

func HandleCommand(c string) {
	if c == "draw" {
		d := cards.Deck{}
		d.Init()
		d.Shuffle()
		n, c := d.Draw()
		if n == 0 {
			e := events.Event{Type: events.EventTypeStatusMessage, Data: "Could not draw card!"}
			events.EventBus <- e
			return
		}
		e := events.Event{Type: events.EventTypeStatusMessage, Data: fmt.Sprintf("Draw card: %s", &c)}
		events.EventBus <- e
	}

}

func (w *CommandWindow) HandleEvent(e *events.Event) {
	//Handle tcell types, we need keypresses
	if e.Type == events.EventTypeTCell {
		et, ok := e.Data.(tcell.Event)
		if !ok {
			return
		}
		switch ev := et.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyBackspace && len(w.typingString) > 0 {
				w.typingString = w.typingString[:len(w.typingString)-1]
				break
			} else if ev.Key() == tcell.KeyEnter {
				e := events.Event{Type: events.EventTypeStatusMessage, Data: fmt.Sprintf("Sending command %s", w.typingString)}
				events.EventBus <- e
				HandleCommand(w.typingString)

				w.typingString = ""
				break
			}
			w.typingString += string(ev.Str())
		}
	}
}
