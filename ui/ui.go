package ui

import "github.com/gdamore/tcell"

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
}

// View is a drawable view on the screen
type View interface {
	Input(*tcell.EventKey)
	Draw(s tcell.Screen)
	Cursor() (int, int)
	Status() string
}

// NewWindow creates and returns a UI window
func NewWindow(x1, y1, x2, y2 int) *Window {
	return &Window{x1, y1, x2, y2}
}

func (w *Window) Size() (int, int) {
	return w.x2 - w.x1, w.y2 - w.y1
}

// Box returns the x1, y1, x2, and y2 of the window
func (w *Window) Box() (int, int, int, int) {
	return w.x1, w.y1, w.x2, w.y2
}

// SetBox updates the bounding box of the window
func (w *Window) SetBox(x1, y1, x2, y2 int) {
	w.x1, w.y1, w.x2, w.y2 = x1, y1, x2, y2
}

// ScrtoWin converts a screen coordinate to window space
func (w *Window) ScrtoWin(x, y int) (int, int) {
	return x - w.x1, y - w.y1
}

// WintoScr converts a window coordinate to screen space
func (w *Window) WintoScr(x, y int) (int, int) {
	return x + w.x1, y + w.y1
}
