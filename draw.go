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

// DrawBody draws the body for the application
func DrawBody() {
	colWidth := 12
	numCols := term.w / colWidth
	light := tcell.ColorWhite
	dark := tcell.ColorBlack
	for y := 1; y < term.h-1; y++ {
		x := 0
		for n := 0; n < numCols; n++ {
			bg := light
			if y%2 == 0 {
				bg = dark
			}
			if n%2 == 0 {
				if bg == light {
					bg = dark
				} else {
					bg = light
				}
			}
			drawBox(&Box{x: x, y: y, w: colWidth, h: 1, bg: bg})
			x += colWidth
		}
	}
}

// DrawFooter draws the footer for the application
func DrawFooter() {
	drawBox(&Box{x: 0, y: term.h - 1, w: term.w, h: 1, bg: tcell.ColorGreen})
}
