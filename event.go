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
		if !cursor.IsVisible() {
			col := cursor.Col
			if resizeH {
				col = grid.ColOff + grid.ColLim - 1
			}
			row := cursor.Row
			if resizeV {
				row = grid.RowOff + grid.RowLim - 1
			}
			cursor.MoveTo(row, col)
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
			cursor.MoveDown()
		case tcell.KeyUp:
			cursor.MoveUp()
		case tcell.KeyLeft:
			cursor.MoveLeft()
		case tcell.KeyRight:
			cursor.MoveRight()
		case tcell.KeyCtrlC:
			quit()
		case tcell.KeyRune:
			switch event.Rune() {
			case 'j':
				cursor.MoveDown()
			case 'k':
				cursor.MoveUp()
			case 'h':
				cursor.MoveLeft()
			case 'l':
				cursor.MoveRight()
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
