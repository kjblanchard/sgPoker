package main

import (
	"log"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/kjblanchard/sgPoker/events"
	"github.com/kjblanchard/sgPoker/ui"
)

func main() {
	defStyle := tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()
	s.SetCursorStyle(tcell.CursorStyleSteadyBar, color.Green)
	x, y := s.Size()
	ui.WindowResizeEvent(x, y)

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()
	events.EventBus = make(chan events.Event, 50)

	//Create and initialize all windows
	var commandW ui.CommandWindow
	commandW.Screen = s
	commandW.Initialize()
	var messageW ui.MessageWindow
	messageW.Screen = s
	messageW.Initialize()

	var playW ui.PlayAreaWindow
	playW.Screen = s
	playW.Initialize()
	var chatW ui.ChatWindow
	chatW.Screen = s
	chatW.Initialize()
	//Handle buildin tcell events, wrap then in a event
	go func() {
		for {
			ev := <-s.EventQ()
			e := events.Event{Type: events.EventTypeTCell, Data: ev}
			events.EventBus <- e
		}
	}()
	//Handle networking messages
	for {
		s.Show() //Update screen
		//Wait on events from our event bus, could be a builtin event, network or from any of our windows sending events
		ev := <-events.EventBus
		if func() bool {
			evc, ok := ev.Data.(tcell.Event)
			if !ok {
				return false
			}
			// Handle if we should quit, resize event, or  something else.
			switch ev := evc.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					return true
				} else if ev.Key() == tcell.KeyCtrlL {
					s.Sync()
				}
			}
			return false

		}() {
			return
		}
		//Pass Event to all Windows to perform their update
		commandW.HandleEvent(&ev)
		messageW.HandleEvent(&ev)
		playW.HandleEvent(&ev)
		chatW.HandleEvent(&ev)
		//Draw all windows
		commandW.Draw()
		messageW.Draw()
		playW.Draw()
		chatW.Draw()
	}
}
