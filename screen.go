package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

var screen tcell.Screen

// StartScreen sets up the screen for the application
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

	initScreen()
}

func initScreen() {
	style := tcell.StyleDefault.
		Background(tcell.ColorBlue).
		Foreground(tcell.ColorReset)
	screen.SetStyle(style)

	screen.Clear()
}

// StopScreen cleans up the screen to return control to the terminal
func StopScreen() {
	screen.Fini()
}
