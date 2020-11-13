package main

import (
	"github.com/gdamore/tcell/v2"
)

func StartEngine() {
	resizeEventChan := make(chan *tcell.EventResize)
	keyEventChan := make(chan *tcell.EventKey)
	quitChan := make(chan bool)

	go startEventLoop(resizeEventChan, keyEventChan)

	go handleResize(resizeEventChan)
	go handleKey(keyEventChan, quitChan)

	<-quitChan
}

func startEventLoop(resizeEventChan chan *tcell.EventResize, keyEventChan chan *tcell.EventKey) {
	for {
		event := screen.PollEvent()

		switch event := event.(type) {
		case *tcell.EventResize:
			resizeEventChan <- event
		case *tcell.EventKey:
			keyEventChan <- event
		}
	}
}
