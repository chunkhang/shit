package main

import (
	"github.com/gdamore/tcell/v2"
)

// HandleResize handles the terminal resize event
func HandleResize() {
	for {
		event := <-channel.resize
		width, height := event.Size()
		term.w = width
		term.h = height
		screen.Clear()
		RefreshScreen()
	}
}

// HandleKey handles the key press event
func HandleKey() {
	for {
		event := <-channel.key
		switch event.Key() {
		case tcell.KeyCtrlC:
			quit()
		case tcell.KeyRune:
			switch event.Rune() {
			case 'j':
				grid.cursor.moveDown()
			case 'k':
				grid.cursor.moveUp()
			case 'h':
				grid.cursor.moveLeft()
			case 'l':
				grid.cursor.moveRight()
			case 'q':
				quit()
			}
		}
		RefreshScreen()
	}
}

func quit() {
	channel.quit <- true
}
