package main

import (
	"github.com/gdamore/tcell/v2"
)

// HandleResize handles the terminal resize event
func HandleResize() {
	for {
		event := <-channel.resize
		w, h := event.Size()

		// Determine resize direction
		resizeH := term.w != w
		resizeV := term.h != h

		// Update screen with new terminal size
		term.w = w
		term.h = h
		RefreshScreen()

		// Reposition cursor if it is gone
		if !grid.cursor.IsVisible() {
			col := grid.cursor.col
			if resizeH {
				col = grid.colOff + grid.colLim - 1
			}
			row := grid.cursor.row
			if resizeV {
				row = grid.rowOff + grid.rowLim - 1
			}
			grid.cursor.MoveTo(row, col)
			RefreshScreen()
		}
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
					grid.cursor.value = "appwlqpeqwlpe"
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
