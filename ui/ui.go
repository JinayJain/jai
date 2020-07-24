package ui

import (
	"github.com/gdamore/tcell"
)

/*

Window system:

- each view is isolated into a window
- views should provide their coordinates in window coordinates

when drawing to the screen:
-

*/

// Window represents a view in the UI
type Window struct {
	x1, y1 int
	x2, y2 int
	View   View
}

// View is a drawable widget used by Window
type View interface {
	// Draw gives coordinates in window space, expecting a rune or false bool
	Draw(int, int) (rune, bool)

	// Cursor returns the window coordinates of where to draw the cursor (when in focus)
	Cursor() (int, int)
}

// NewWindow creates and returns a UI window
func NewWindow(v View, x1, y1, x2, y2 int) *Window {
	return &Window{x1, y1, x2, y2, v}
}

// Box returns the x1, y1, x2, and y2 of the window
func (w *Window) Box() (int, int, int, int) {
	return w.x1, w.y1, w.x2, w.y2
}

// SetBox updates the bounding box of the window
func (w *Window) SetBox(x1, y1, x2, y2 int) {
	w.x1, w.y1, w.x2, w.y2 = x1, y1, x2, y2
}

// Draw renders the current window view onto the screen
func (w *Window) Draw(s tcell.Screen, focused bool) {

	if focused {
		s.ShowCursor(w.WintoScr(w.View.Cursor()))
	}

	for i := w.y1; i <= w.y2; i++ {
		for j := w.x1; j <= w.x2; j++ {
			r, valid := w.View.Draw(w.ScrtoWin(j, i))
			if valid {
				s.SetContent(j, i, r, nil, tcell.StyleDefault)
			}
		}
	}
}

// ScrtoWin converts a screen coordinate to window space
func (w *Window) ScrtoWin(x, y int) (int, int) {
	return x - w.x1, y - w.y1
}

// WintoScr converts a window coordinate to screen space
func (w *Window) WintoScr(x, y int) (int, int) {
	return x + w.x1, y + w.y1
}
