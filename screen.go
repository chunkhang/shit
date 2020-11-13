package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

var screen tcell.Screen

func startScreen() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	screen = s

	err = screen.Init()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	style := tcell.StyleDefault.
		Background(tcell.ColorBlue).
		Foreground(tcell.ColorReset)
	screen.SetStyle(style)

	screen.Clear()
}

func stopScreen() {
	screen.Fini()
}
