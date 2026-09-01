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
	busChannel := make(chan events.Event)

	//Create and initialize all windows
	var commandW ui.CommandWindow
	commandW.Screen = s
	commandW.Initialize()

	for {
		s.Show() //Update screen
		//Handle buildin tcell events, wrap then in a event
		go func() {
			ev := <-s.EventQ()
			e := events.Event{Type: 0, Data: ev}
			busChannel <- e
		}()
		//Handle networking messages
		//Wait on events from our event bus, could be a builtin event, network or from any of our windows sending events
		ev := <-busChannel
		evc, ok := ev.Data.(tcell.Event)
		if !ok {
			log.Print("Not a correct type!")
			continue
		}
		// Handle if we should quit, resize event, or  something else.
		switch ev := evc.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			}
		}
		//Pass Event to all Windows to perform their update
		commandW.HandleEvent(&ev)
		//Draw all windows
		commandW.Draw()
	}
}
