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
			case 'q':
				channel.quit <- true
			}
		}
		RefreshScreen()
	}
}
