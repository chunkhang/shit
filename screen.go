package main

import (
	"fmt"
	"strconv"

	cv "github.com/chunkhang/shit/canvas"
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
	screen.Clear()
	drawHeader()
	drawBody()
	drawFooter()
	screen.Show()
}

const (
	headerHeight = 1
	lineHeight   = 1
	lineWidth    = 2
	cellHeight   = 1
	cellWidth    = 9
	footerHeight = 2
)

func drawHeader() {
	text := fmt.Sprintf("%s%s", ColLabel(cursor.Col), RowLabel(cursor.Row))
	canvas.NewBox(0, 0, term.w, headerHeight).
		Text(fmt.Sprintf("%s %s", text, cursor.Value)).
		Draw()
}

func drawBody() {
	yStart := headerHeight
	yEnd := term.h - footerHeight

	xStart := 0
	xEnd := term.w

	rowStart := grid.RowOff
	rowLim := yEnd - yStart - (2 * cellHeight)
	rowEnd := Min(rowStart+rowLim, grid.RowTotal)
	grid.RowLim = rowLim

	// Row index width is determined by number of digits in the largest row number
	// Add 2 to this number for padding
	rowIndexWidth := len(strconv.Itoa(rowEnd)) + 2
	colStart := grid.ColOff
	colLim := (xEnd - xStart - rowIndexWidth) / cellWidth
	colEnd := Min(colStart+colLim, grid.ColTotal)
	grid.ColLim = colLim

	// Column index
	for col := colStart; col < colEnd; col++ {
		x := xStart + lineWidth + (col-colStart)*cellWidth + rowIndexWidth
		y := yStart
		canvas.NewBox(x, y, cellWidth, cellHeight).
			Reverse(col == cursor.Col).
			Text(ColLabel(col)).
			AlignCenter().
			Draw()
	}

	xStartHLine := xStart + rowIndexWidth
	yHLine := yStart + cellHeight

	// Horizontal line
	for x := xStartHLine; x < xEnd; x++ {
		canvas.NewPoint(x, yHLine).
			Char(tcell.RuneHLine).
			Draw()
	}

	// Row index
	for row := rowStart; row < rowEnd; row++ {
		x := xStart
		y := yStart + (row-rowStart+1)*cellHeight + lineHeight
		canvas.NewBox(x, y, rowIndexWidth, cellHeight).
			Reverse(row == cursor.Row).
			Text(RowLabel(row)).
			AlignRight().
			Pad(1).
			Draw()
	}

	yStartVLine := yStart + cellHeight
	xVLine := xStart + rowIndexWidth

	// Vertical line
	for y := yStartVLine; y < yEnd; y++ {
		canvas.NewPoint(xVLine, y).
			Char(tcell.RuneVLine).
			Draw()
	}

	// Line intersection
	canvas.NewPoint(xVLine, yHLine).
		Char(tcell.RuneULCorner).
		Draw()

	// Cells
	for row := rowStart; row < rowEnd; row++ {
		for col := colStart; col < colEnd; col++ {
			cell := grid.GetCell(row, col)
			x := xStart + lineWidth + (col-colStart)*cellWidth + rowIndexWidth
			y := yStart + lineHeight + (row-rowStart+1)*cellHeight
			canvas.NewBox(x, y, cellWidth, cellHeight).
				Reverse(row == cursor.Row && col == cursor.Col).
				Text(cell.Value).
				PadRight(1).
				Draw()
		}
	}
}

func drawFooter() {
	canvas.NewBox(0, term.h-footerHeight, term.w, footerHeight).
		Text(file.Name).
		Draw()
}
