package main

import (
	"fmt"
	"strconv"

	cv "github.com/chunkhang/xr/canvas"
	"github.com/gdamore/tcell/v2"
)

var screen tcell.Screen
var canvas *cv.Canvas

// StartScreen sets up the screen for the application
func StartScreen() (err error) {
	screen, err = tcell.NewScreen()
	if err != nil {
		return
	}

	err = screen.Init()
	if err != nil {
		return
	}

	reset := tcell.ColorReset
	style := tcell.StyleDefault.Background(reset).Foreground(reset)
	screen.SetStyle(style)

	screen.Clear()

	canvas = &cv.Canvas{
		Screen: screen,
	}

	return
}

// StopScreen cleans up the screen to return control to the terminal
func StopScreen() {
	screen.Fini()
}

// RefreshScreen redraws the screen with the latest application state
func RefreshScreen() {
	DrawHeader()
	DrawBody()
	screen.Show()
}

const (
	headerHeight = 1
)

// DrawHeader draws the header for the application
func DrawHeader() {
	canvas.NewBox(0, 0, term.w, headerHeight).
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorSilver).
		Text(fmt.Sprintf("%s %s", grid.cursor.pos, grid.cursor.value)).
		Draw()
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
		x := xStart + col*cellWidth + rowIndexWidth
		y := yStart
		canvas.NewBox(x, y, cellWidth, cellHeight).
			Background(bg).
			Foreground(fg).
			Text(strconv.Itoa(col)).
			AlignCenter().
			Draw()
	}

	// Row index
	for row := 0; row < rowNum; row++ {
		bg, fg := tcell.ColorBlack, tcell.ColorSilver
		if row == grid.cursor.pos.row {
			bg, fg = fg, bg
		}
		x := xStart
		y := yStart + row*cellHeight + cellHeight
		canvas.NewBox(x, y, rowIndexWidth, cellHeight).
			Background(bg).
			Foreground(fg).
			Text(strconv.Itoa(row)).
			AlignRight().
			Pad(1).
			Draw()
	}

	// Cells
	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			bg, fg := tcell.ColorBlack, tcell.ColorSilver
			if row == grid.cursor.pos.row && col == grid.cursor.pos.col {
				bg, fg = fg, bg
			}
			cell := grid.GetCell(&Pos{row: row, col: col})
			x := xStart + col*cellWidth + rowIndexWidth
			y := yStart + row*cellHeight + cellHeight
			canvas.NewBox(x, y, cellWidth, cellHeight).
				Background(bg).
				Foreground(fg).
				Text(cell.value).
				PadRight(1).
				Draw()
		}
	}
}
