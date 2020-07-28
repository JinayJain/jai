package ui

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/jinayjain/jai/file"
)

/*

Editor:

- maintains the internal buffer as a slice of lines, each line being a slice of cells
- should the editor maintain the current origin/offset of the view
	- also cursor positioning


Methods:

[x] add a character
[x] delete a character
[x] add a line
[x] delete a line
[ ] split a line
[ ] move the window
[x] move the cursor

*/

// Mode represents a possible editor mode
type Mode int

// The possible modes the editor can be in
const (
	Edit Mode = iota
	Write
)

func (m Mode) String() string {
	switch m {
	case Edit:
		return "Edit"
	case Write:
		return "Write"
	default:
		return fmt.Sprintf("%d", int(m))
	}
}

// Editor defines an editor window for editing a single file
type Editor struct {
	Window *Window
	Buffer [][]rune
	Path   string
	Mode   Mode
	cx, cy int // cursor's x and y position in the buffer
	ox, oy int // x and y offset of the viewing window
	cnt    int
}

// NewEditor creates and returns an editor instance
func NewEditor(path string, x1, y1, x2, y2 int) *Editor {
	e := Editor{}
	e.Window = NewWindow(x1, y1, x2, y2)
	e.Load(path)

	return &e
}

// Load reads the given file into the buffer.
func (e *Editor) Load(path string) {
	e.Path = path
	if path != "" {
		f, err := file.Read(path)

		if err == nil {
			e.Buffer = append(e.Buffer, f...)
		}
	}

	if len(e.Buffer) == 0 {
		e.Buffer = make([][]rune, 1)
	}
}

// Save is used to write the buffer to a path
func (e *Editor) Save(path string) {
	err := file.Write(path, e.Buffer)
	if err != nil {
		panic(err)
	}
}

// Cursor returns the window position of the cursor in the buffer
func (e *Editor) Cursor() (int, int) {
	return e.Window.WintoScr(e.BuftoWin(e.cx, e.cy))
}

// MoveCursor moves the cursor within boundaries, shifting the offset if necessary
func (e *Editor) MoveCursor(dx, dy int) {

	nx, ny := e.cx+dx, e.cy+dy

	if nx < 0 {
		nx = 0
	}
	if ny < 0 {
		ny = 0
	}
	if ny >= len(e.Buffer) {
		ny = len(e.Buffer) - 1
	}
	if nx > len(e.Buffer[ny]) {
		nx = len(e.Buffer[ny])
	}

	wx, wy := e.BuftoWin(nx, ny)
	if wy < e.Window.y1 {
		e.oy -= e.Window.y1 - wy
	} else if wy > e.Window.y2 {
		e.oy += wy - e.Window.y2
	}

	if wx < e.Window.x1 {
		e.ox -= e.Window.x1 - wx
	} else if wx > e.Window.x2 {
		e.ox += wx - e.Window.x2
	}

	e.cx, e.cy = nx, ny
}

// Input performs an editor action based on user input
func (e *Editor) Input(ev *tcell.EventKey) {
	if ev.Key() == tcell.KeyCtrlS {
		e.Save(e.Path)
		return
	}

	switch e.Mode {
	case Write:
		e.inputInsert(ev)
	case Edit:
		e.inputEdit(ev)
	}
}

func (e *Editor) inputInsert(ev *tcell.EventKey) {

	key := ev.Key()

	switch key {
	case tcell.KeyRune:
		ch := ev.Rune()

		e.Buffer[e.cy] = append(e.Buffer[e.cy], 0)
		copy(e.Buffer[e.cy][e.cx+1:], e.Buffer[e.cy][e.cx:])
		e.Buffer[e.cy][e.cx] = ch

		e.MoveCursor(1, 0)

	case tcell.KeyBackspace:
	case tcell.KeyBackspace2:
		if len(e.Buffer[e.cy]) > 0 && e.cx > 0 {
			e.Buffer[e.cy] = append(e.Buffer[e.cy][:e.cx-1], e.Buffer[e.cy][e.cx:]...)
			e.MoveCursor(-1, 0)
		}

	case tcell.KeyEnter:
		e.Buffer = append(e.Buffer, []rune{})
		e.MoveCursor(0, 1)
	case tcell.KeyEscape:
		e.Mode = Edit
	}
}

func (e *Editor) inputEdit(ev *tcell.EventKey) {
	key := ev.Key()

	switch key {
	case tcell.KeyRune:
		switch ev.Rune() {
		case 'h':
			e.MoveCursor(-1, 0)
		case 'j':
			e.MoveCursor(0, 1)
		case 'k':
			e.MoveCursor(0, -1)
		case 'l':
			e.MoveCursor(1, 0)
		case 'd':
			switch {
			case len(e.Buffer) == 0:
			case len(e.Buffer) == 1:
				e.Buffer = make([][]rune, 1)
			default:
				e.Buffer = append(e.Buffer[:e.cy], e.Buffer[e.cy+1:]...)
			}
			e.MoveCursor(0, 0) // hacky way of making sure the cursor is within bounds ;)
		case 'o':
			e.Buffer = append(e.Buffer, []rune{})
			copy(e.Buffer[e.cy+2:], e.Buffer[e.cy+1:])
			e.Buffer[e.cy+1] = []rune{}
			e.MoveCursor(0, 1)
			e.Mode = Write

		case 'i':
			e.Mode = Write
		}
	}
}

// Draw renders the editor to the screen
func (e *Editor) Draw(s tcell.Screen) {
	x1, y1, x2, y2 := e.Window.Box()

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			r, valid := e.getRune(e.Window.ScrtoWin(i, j))
			if valid {
				s.SetContent(i, j, r, nil, tcell.StyleDefault)
			}
		}
	}
}

// Status returns the status string for display in the status bar
func (e *Editor) Status() (status string) {
	if e.Path != "" {
		status += fmt.Sprintf("%s ", e.Path)
	}

	status += fmt.Sprintf("[%s Mode]", e.Mode.String())
	return status
}

func (e *Editor) getRune(x, y int) (rune, bool) {
	x, y = e.WintoBuf(x, y)

	if y >= len(e.Buffer) {
		return 0, false
	}

	if x >= len(e.Buffer[y]) {
		return 0, false
	}

	return e.Buffer[y][x], true
}

// BuftoWin converts a buffer coordinate to window space
func (e *Editor) BuftoWin(x, y int) (int, int) {
	return x - e.ox, y - e.oy
}

// WintoBuf converts a window coordinate to buffer space
func (e *Editor) WintoBuf(x, y int) (int, int) {
	return x + e.ox, y + e.oy
}
