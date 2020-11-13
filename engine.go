package main

import (
	"github.com/gdamore/tcell/v2"
)

// StartEngine starts the event loop for the application
// All event handlers are registered as well
// The engine will stop after receiving the quit signal
func StartEngine() {
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

	go HandleResize()
	go HandleKey()

	<-channel.quit
}
