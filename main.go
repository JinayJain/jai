package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/jinayjain/jai/ui"
)

func main() {

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, err := tcell.NewScreen()
	check(err)

	err = s.Init()
	check(err)

	w, h := s.Size()

	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	m := ui.NewManager(path, w, h)

main:
	for {
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyCtrlQ:
				break main
			default:
				m.Input(ev)
			}
		case *tcell.EventResize:
			s.Sync()
		}

		s.Clear()

		m.Draw(s)

		s.Show()
	}
	s.Fini()
}

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
