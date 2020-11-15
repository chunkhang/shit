package main

import (
	"fmt"
	"strconv"

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

const (
	headerHeight = 1
)

// DrawHeader draws the header for the application
func DrawHeader() {
	drawBox(&Box{
		x:    0,
		y:    0,
		w:    term.w,
		h:    headerHeight,
		bg:   tcell.ColorBlack,
		fg:   tcell.ColorSilver,
		text: fmt.Sprintf("%s %s", grid.cursor.pos, grid.cursor.value),
	})
}

const (
	cellHeight = 1
	cellWidth  = 9
)

// DrawBody draws the body for the application
func DrawBody() {
	yStart := headerHeight
	yEnd := term.h

	xStart := 0
	xEnd := term.w

	rowNum := yEnd - yStart - cellHeight

	// Row index width is determined by number of digits in the largest row number
	// Add 2 to this number for padding
	rowIndexWidth := len(strconv.Itoa(rowNum)) + 2
	colNum := (xEnd - xStart - rowIndexWidth) / cellWidth

	// Column index
	for col := 0; col < colNum; col++ {
		bg, fg := tcell.ColorBlack, tcell.ColorSilver
		if col == grid.cursor.pos.col {
			bg, fg = fg, bg
		}
		drawBox(&Box{
			x:     xStart + col*cellWidth + rowIndexWidth,
			y:     yStart,
			w:     cellWidth,
			h:     cellHeight,
			bg:    bg,
			fg:    fg,
			text:  strconv.Itoa(col),
			align: alignCenter,
		})
	}

	// Row index
	for row := 0; row < rowNum; row++ {
		bg, fg := tcell.ColorBlack, tcell.ColorSilver
		if row == grid.cursor.pos.row {
			bg, fg = fg, bg
		}
		drawBox(&Box{
			x:     xStart,
			y:     yStart + row*cellHeight + cellHeight,
			w:     rowIndexWidth,
			h:     cellHeight,
			bg:    bg,
			fg:    fg,
			text:  fmt.Sprintf("%d ", row),
			align: alignRight,
		})
	}

	// Cells
	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			bg, fg := tcell.ColorBlack, tcell.ColorSilver
			if row == grid.cursor.pos.row && col == grid.cursor.pos.col {
				bg, fg = fg, bg
			}
			cell := grid.GetCell(&Pos{row: row, col: col})
			drawBox(&Box{
				x:    xStart + col*cellWidth + rowIndexWidth,
				y:    yStart + row*cellHeight + cellHeight,
				w:    cellWidth,
				h:    cellHeight,
				bg:   bg,
				fg:   fg,
				text: cell.value,
			})
		}
	}
}
