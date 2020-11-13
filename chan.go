package main

import (
	"github.com/gdamore/tcell/v2"
)

type chanGroup struct {
	resize chan *tcell.EventResize
	key    chan *tcell.EventKey
	quit   chan bool
}

var channel = &chanGroup{
	resize: make(chan *tcell.EventResize),
	key:    make(chan *tcell.EventKey),
	quit:   make(chan bool),
}
