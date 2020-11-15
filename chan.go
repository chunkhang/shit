package main

import (
	"github.com/gdamore/tcell/v2"
)

// Chan is all the channels used by the application
type Chan struct {
	resize chan *tcell.EventResize
	key    chan *tcell.EventKey
	quit   chan bool
}

var channel = &Chan{
	resize: make(chan *tcell.EventResize),
	key:    make(chan *tcell.EventKey),
	quit:   make(chan bool),
}
