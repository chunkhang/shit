package main

import (
	"github.com/gdamore/tcell/v2"
)

// Point is the smallest drawable point
type Point struct {
	x    int
	y    int
	bg   tcell.Color
	fg   tcell.Color
	char rune
}

func drawPoint(point *Point) {
	style := tcell.StyleDefault.Background(point.bg).Foreground(point.fg)
	screen.SetContent(point.x, point.y, point.char, nil, style)
}

// Box is a drawable rectangular area
type Box struct {
	x    int
	y    int
	w    int
	h    int
	bg   tcell.Color
	fg   tcell.Color
	text string
}

// drawBox draws a box with text on the screen
func drawBox(box *Box) {
	i := 0
	chars := []rune(box.text)
	for y := box.y; y < box.y+box.h; y++ {
		for x := box.x; x < box.x+box.w; x++ {
			point := &Point{x: x, y: y, bg: box.bg, fg: box.fg}
			if i < len(chars) {
				point.char = chars[i]
				i++
			}
			drawPoint(point)
		}
	}
}

// DrawHeader draws the header for the application
func DrawHeader() {
	drawBox(&Box{x: 0, y: 0, w: term.w, h: 1, bg: tcell.ColorBlue})
}

// DrawFooter draws the footer for the application
func DrawFooter() {
	drawBox(&Box{x: 0, y: term.h - 1, w: term.w, h: 1, bg: tcell.ColorGreen})
}
