package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
)

func StartEngine() {
	for {
		screen.Show()
		event := screen.PollEvent()

		switch event := event.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyEscape:
				log.Println("esc")
			case tcell.KeyCtrlC:
				log.Println("ctrl-c")
			case tcell.KeyRune:
				switch event.Rune() {
				case 'q':
					log.Println("q")
					return
				}
			}
		}
	}
}
