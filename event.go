package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

// HandleResize handles the terminal resize event
func HandleResize() {
	for {
		event := <-channel.resize
		width, height := event.Size()
		log.Printf("width, height = %+v, %+v\n", width, height)
		screen.Sync()
	}
}

// HandleKey handles the key press event
func HandleKey() {
loop:
	for {
		event := <-channel.key
		log.Printf("event.Name() = %+v\n", event.Name())
		switch event.Key() {
		case tcell.KeyEscape:
			break loop
		case tcell.KeyCtrlC:
			break loop
		case tcell.KeyRune:
			switch event.Rune() {
			case 'b':
				style := tcell.StyleDefault.
					Background(tcell.ColorBlue)
				screen.SetStyle(style)
				screen.Clear()
			case 'g':
				style := tcell.StyleDefault.
					Background(tcell.ColorGreen)
				screen.SetStyle(style)
				screen.Clear()
			case 'r':
				style := tcell.StyleDefault.
					Background(tcell.ColorRed)
				screen.SetStyle(style)
				screen.Clear()
			case 'q':
				break loop
			}
		}
		screen.Show()
	}
	channel.quit <- true
}
