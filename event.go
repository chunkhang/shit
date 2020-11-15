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
		case tcell.KeyDown:
			grid.cursor.MoveDown()
		case tcell.KeyUp:
			grid.cursor.MoveUp()
		case tcell.KeyLeft:
			grid.cursor.MoveLeft()
		case tcell.KeyRight:
			grid.cursor.MoveRight()
		case tcell.KeyCtrlC:
			quit()
		case tcell.KeyRune:
			switch event.Rune() {
			case 'j':
				grid.cursor.MoveDown()
			case 'k':
				grid.cursor.MoveUp()
			case 'h':
				grid.cursor.MoveLeft()
			case 'l':
				grid.cursor.MoveRight()
			case ' ':
				if grid.cursor.value == "" {
					grid.cursor.value = "*"
				} else {
					grid.cursor.value = ""
				}
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
