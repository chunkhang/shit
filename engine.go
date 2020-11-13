package main

import (
	"github.com/gdamore/tcell/v2"
)

func startEngine() {
	// Event loop
	go func() {
		for {
			event := screen.PollEvent()
			switch event := event.(type) {
			case *tcell.EventResize:
				channel.resize <- event
			case *tcell.EventKey:
				channel.key <- event
			}
		}
	}()

	// Handle events
	go handleResize()
	go handleKey()

	// Stop engine upon receiving quit signal
	<-channel.quit
}
