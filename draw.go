package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Drawable interface {
	Draw()
}

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

// Align is direction of text alignment
type Align int

const (
	alignLeft Align = iota
	alignRight
	alignCenter
)

// Box is a drawable rectangular area
type Box struct {
	x     int
	y     int
	w     int
	h     int
	bg    tcell.Color
	fg    tcell.Color
	text  string
	align Align
}

func drawBox(box *Box) {
	yStart := box.y
	yEnd := box.y + box.h

	xStart := box.x
	xEnd := box.x + box.w
	xLength := xEnd - xStart

	chars := []rune(box.text)
	charTotal := len(chars)
	i := 0

	for y := yStart; y < yEnd; y++ {
		// Calculate remaining characters for this row
		// This information is used for text alignment
		charOffset := 0
		charRemain := charTotal - i
		if charRemain < xLength {
			if box.align == alignRight {
				charOffset = xLength - charRemain
			} else if box.align == alignCenter {
				charOffset = (xLength - charRemain) / 2
			}
		}
		for x := xStart; x < xEnd; x++ {
			point := &Point{x: x, y: y, bg: box.bg, fg: box.fg}
			offset := x - xStart
			if i < charTotal && offset >= charOffset {
				point.char = chars[i]
				i++
			}
			drawPoint(point)
		}
	}
}

// DrawHeader draws the header for the application
func DrawHeader() {
	text := fmt.Sprintf("%s %s", grid.cursor.pos, grid.cursor.value)
	drawBox(&Box{x: 0, y: 0, w: term.w, h: 1, bg: tcell.ColorBlack, fg: tcell.ColorSilver, text: text})
}

// DrawBody draws the body for the application
func DrawBody() {
	colWidth := 12
	numCols := term.w / colWidth
	// Row headings
	for n := 1; n < numCols; n++ {
		x := n * colWidth
		col := n - 1
		bg := tcell.ColorBlack
		fg := tcell.ColorSilver
		// Highlight cursor column
		if col == grid.cursor.pos.col {
			bg, fg = fg, bg
		}
		box := &Box{x: x, y: 1, w: colWidth, h: 1, bg: bg, fg: fg, text: fmt.Sprintf("%d", col), align: alignCenter}
		drawBox(box)
	}
	// Column headings
	for y := 2; y < term.h-1; y++ {
		row := y - 2
		bg := tcell.ColorBlack
		fg := tcell.ColorSilver
		// Highlight cursor row
		if row == grid.cursor.pos.row {
			bg, fg = fg, bg
		}
		box := &Box{x: 0, y: y, w: colWidth, h: 1, bg: bg, fg: fg, text: fmt.Sprintf("%d", row), align: alignRight}
		drawBox(box)
	}
	// Cells
	for y := 2; y < term.h-1; y++ {
		x := 0
		for n := 0; n < numCols-1; n++ {
			row := y - 2
			col := n
			bg := tcell.ColorBlack
			fg := tcell.ColorSilver
			// Highlight cursor position
			if row == grid.cursor.pos.row && col == grid.cursor.pos.col {
				bg, fg = fg, bg
			}
			cell := grid.GetCell(&Pos{row: row, col: col})
			box := &Box{x: x + colWidth, y: y, w: colWidth, h: 1, bg: bg, fg: fg, text: cell.value}
			drawBox(box)
			x += colWidth
		}
	}
}

// DrawFooter draws the footer for the application
func DrawFooter() {
	drawBox(&Box{x: 0, y: term.h - 1, w: term.w, h: 1})
}
