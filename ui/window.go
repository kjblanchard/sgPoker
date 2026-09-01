package ui

import (
	"github.com/gdamore/tcell/v3"
	"github.com/kjblanchard/sgPoker/events"
)

var (
	screenX int
	screenY int
)

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if text == "" {
		return
	}
	row := y1
	col := x1
	var width int
	for text != "" {
		text, width = s.Put(col, row, text, style)
		col += width
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
		if width == 0 {
			// incomplete grapheme at end of string
			break
		}
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.Put(col, row, " ", style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.Put(col, y1, string(tcell.RuneHLine), style)
		s.Put(col, y2, string(tcell.RuneHLine), style)
	}
	//Handle Drawing title
	if text != "" {
		halfWord := len(text) / 2
		half := x2 / 2
		titleLoc := half - halfWord
		for i, l := range text {
			s.Put(titleLoc+i, y1, string(l), style)

		}

	}
	for row := y1 + 1; row < y2; row++ {
		s.Put(x1, row, string(tcell.RuneVLine), style)
		s.Put(x2, row, string(tcell.RuneVLine), style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.Put(x1, y1, string(tcell.RuneULCorner), style)
		s.Put(x2, y1, string(tcell.RuneURCorner), style)
		s.Put(x1, y2, string(tcell.RuneLLCorner), style)
		s.Put(x2, y2, string(tcell.RuneLRCorner), style)
	}

	// drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

func WindowResizeEvent(x int, y int) {
	screenX = x
	screenY = y
}

// Change the size of the window based on the screen size.  Useful for making "dynamic" screens.
// X and y are percents, xo and yo are pixels for the "borders" around it that it will subtract from the percentage called x and y offsets.
func SetWindowSizeByPercent(w *WindowBase, x float32, y float32, xo int, yo int) {
	//determine size based on screen
	w.W = int(float32(screenX)*x) - xo
	w.H = int(float32(screenY)*y) - yo
}

type WindowBase struct {
	X      int
	Y      int
	W      int
	H      int
	Screen tcell.Screen
}

type Window interface {
	Initialize()
	Draw()
	HandleEvent(e *events.Event)
}
