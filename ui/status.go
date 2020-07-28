package ui

import "github.com/gdamore/tcell"

// StatusBar is the status bar shown for each editor
type StatusBar struct {
	Window *Window
	Active View
}

// NewStatusBar creates and returns the default status bar
func NewStatusBar(width, y int) *StatusBar {
	win := NewWindow(0, y, width, y)
	return &StatusBar{Window: win}
}

// Draw renders the status bar to the screen
func (sb *StatusBar) Draw(s tcell.Screen) {
	str := []rune(sb.Active.Status())
	/*
		for centered status bar
		width, _ := sb.Window.Size()
		offset := width/2 - len(str)/2 + 1
	*/
	offset := 0
	style := tcell.StyleDefault.Bold(true).Foreground(tcell.ColorLightGreen)
	for i, c := range str {
		s.SetContent(offset+i, sb.Window.y1, c, nil, style)
	}
}
