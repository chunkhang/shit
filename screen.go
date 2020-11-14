package main

import (
	"github.com/gdamore/tcell/v2"
)

var screen tcell.Screen

// StartScreen sets up the screen for the application
func StartScreen() (err error) {
	screen, err = tcell.NewScreen()
	if err != nil {
		return
	}

	err = screen.Init()
	if err != nil {
		return
	}

	reset := tcell.ColorReset
	style := tcell.StyleDefault.Background(reset).Foreground(reset)
	screen.SetStyle(style)

	screen.Clear()

	return
}

// RefreshScreen redraws the screen with the latest application state
func RefreshScreen() {
	DrawHeader()
	DrawFooter()
	screen.Show()
}

// StopScreen cleans up the screen to return control to the terminal
func StopScreen() {
	screen.Fini()
}
