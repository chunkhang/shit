package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

var screen tcell.Screen

func StartScreen() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	screen = s

	err = screen.Init()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorReset)
	screen.SetStyle(defStyle)

	screen.Clear()
}

func StopScreen() {
	screen.Fini()
}
