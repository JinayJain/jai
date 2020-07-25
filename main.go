package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/jinayjain/jai/editor"
	"github.com/jinayjain/jai/ui"
)

func main() {

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, err := tcell.NewScreen()
	check(err)

	err = s.Init()
	check(err)

	w, h := s.Size()
	fmt.Println(w, h)

	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	ed := editor.NewEditor(path)
	win := ui.NewWindow(ed, 0, 0, w, h)

	fmt.Println(win.Box())

main:
	for {
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyCtrlC:
				break main
			default:
				ed.Input(ev)
			}
		case *tcell.EventResize:
			s.Sync()
		}

		s.Clear()

		win.Draw(s, true)

		s.Show()
	}

	s.Fini()

}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
