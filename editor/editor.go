package editor

import "github.com/gdamore/tcell"

/*

Editor:

- maintains the internal buffer as a slice of lines, each line being a slice of cells
- should the editor maintain the current origin/offset of the view
	- also cursor positioning


Methods:

- add a character
- delete a character
- add a line
- delete a line
- split a line
- move the window
- move the cursor

*/

// Mode represents a possible editor mode
type Mode int

// The possible modes the editor can be in
const (
	Edit Mode = iota
	Insert
)

// Editor defines an editor window for editing a single file
type Editor struct {
	Buffer [][]rune
	Mode   Mode
	cx, cy int // cursor's x and y position in the buffer
	ox, oy int // x and y offset of the viewing window
}

// NewEditor creates and returns an editor instance
func NewEditor() *Editor {
	e := Editor{}
	e.Buffer = append(e.Buffer, []rune("Hello"))
	e.Buffer = append(e.Buffer, []rune("World"))

	return &e
}

// Cursor returns the position of the cursor in the buffer
func (e *Editor) Cursor() (int, int) {
	return e.cx - e.ox, e.cy
}

// MoveCursor moves the cursor within boundaries
func (e *Editor) MoveCursor(dx, dy int) {
	// TODO: If the cursor goes past window bounds, shift the offset

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

	e.cx, e.cy = nx, ny
}

// Input performs an editor action based on user input
func (e *Editor) Input(ev *tcell.EventKey) {
	switch e.Mode {
	case Insert:
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
		e.Buffer[e.cy] = append(e.Buffer[e.cy], ch)
	case tcell.KeyBackspace:
	case tcell.KeyBackspace2:
		if len(e.Buffer[e.cy]) > 0 {
			e.Buffer[e.cy] = e.Buffer[e.cy][:len(e.Buffer[e.cy])-1]
		}
	case tcell.KeyEnter:
		e.Buffer = append(e.Buffer, []rune{})
		e.cy++
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
		case 'i':
			e.Mode = Insert
		}
	}
}

// Draw implements the View interface
func (e *Editor) Draw(x, y int) (rune, bool) {
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
