package ui

import "github.com/gdamore/tcell"

// StatusBar is the status bar shown for each editor
type StatusBar struct {
	Window *Window
	Active View
}

func NewStatusBar(width, y int) *StatusBar {
	win := NewWindow(0, y, width, y)
	return &StatusBar{Window: win}
}

func (sb *StatusBar) Draw(s tcell.Screen) {
	str := []rune(sb.Active.Status())
	/*
		for centered status bar
		width, _ := sb.Window.Size()
		offset := width/2 - len(str)/2 + 1
	*/
	offset := 0
	for i, c := range str {
		s.SetContent(offset+i, sb.Window.y1, c, nil, tcell.StyleDefault.Reverse(true))
	}
}
