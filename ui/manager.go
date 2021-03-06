package ui

import "github.com/gdamore/tcell"

// Manager is the main UI handler for each instance of a program
type Manager struct {
	Status  *StatusBar
	Editor  *Editor // TODO: Add support for multiple editors
	Focused View
}

// NewManager creates an returns the standard initial layout for a jai instance
func NewManager(path string, w, h int) *Manager {
	ed := NewEditor(path, 0, 0, w-1, h-2)
	sb := NewStatusBar(w-1, h-1)
	sb.Active = ed

	return &Manager{Status: sb, Editor: ed, Focused: ed}
}

// Draw renders all current windows the the screen
func (m *Manager) Draw(s tcell.Screen) {
	s.ShowCursor(m.Editor.Cursor())
	m.Editor.Draw(s)
	m.Status.Draw(s)
}

// Input delegates input events to the respective Views
func (m *Manager) Input(ev *tcell.EventKey) {
	m.Editor.Input(ev)
}
