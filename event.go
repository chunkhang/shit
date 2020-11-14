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
			channel.quit <- true
		case tcell.KeyRune:
			switch event.Rune() {
			case 'j':
				grid.row++
			case 'k':
				if grid.row > 0 {
					grid.row--
				}
			case 'h':
				if grid.col > 0 {
					grid.col--
				}
			case 'l':
				grid.col++
			case ' ':
				pos := &Pos{row: grid.row, col: grid.col}
				if grid.hasValue(pos) {
					grid.setValue(pos, "")
				} else {
					grid.setValue(pos, "*")
				}
			case 'q':
				channel.quit <- true
			}
		}
		RefreshScreen()
	}
}
