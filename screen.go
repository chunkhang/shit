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
	drawHeader()
	drawBody()
	drawFooter()
	screen.Show()
}

const (
	headerHeight = 1
	cellHeight   = 1
	cellWidth    = 9
	footerHeight = 2
)

func drawHeader() {
	canvas.NewBox(0, 0, term.w, headerHeight).
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorSilver).
		Text(fmt.Sprintf("%s %s", grid.cursor.pos.Label(), grid.cursor.value)).
		Draw()
}

func drawBody() {
	yStart := headerHeight
	yEnd := term.h - footerHeight

	xStart := 0
	xEnd := term.w

	rowNum := yEnd - yStart - cellHeight

	// Row index width is determined by number of digits in the largest row number
	// Add 2 to this number for padding
	rowIndexWidth := len(strconv.Itoa(rowNum)) + 2
	colNum := (xEnd - xStart - rowIndexWidth) / cellWidth

	lineHeight := 1
	lineWidth := 2

	// Column index
	for col := 0; col < colNum; col++ {
		bg, fg := tcell.ColorBlack, tcell.ColorSilver
		if col == grid.cursor.pos.col {
			bg, fg = fg, bg
		}
		x := xStart + lineWidth + col*cellWidth + rowIndexWidth
		y := yStart
		canvas.NewBox(x, y, cellWidth, cellHeight).
			Background(bg).
			Foreground(fg).
			Text(ColLabel(col)).
			AlignCenter().
			Draw()
	}

	xStartHLine := xStart + rowIndexWidth
	xEndHLine := xStart + lineWidth + rowIndexWidth + colNum*cellWidth
	yHLine := yStart + cellHeight

	// Horizontal line
	for x := xStartHLine; x < xEndHLine; x++ {
		canvas.NewPoint(x, yHLine).
			Foreground(tcell.ColorWhite).
			Char(tcell.RuneHLine).
			Draw()
	}

	// Row index
	for row := 0; row < rowNum; row++ {
		bg, fg := tcell.ColorBlack, tcell.ColorSilver
		if row == grid.cursor.pos.row {
			bg, fg = fg, bg
		}
		x := xStart
		y := yStart + row*cellHeight + cellHeight + lineHeight
		canvas.NewBox(x, y, rowIndexWidth, cellHeight).
			Background(bg).
			Foreground(fg).
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
			Foreground(tcell.ColorWhite).
			Char(tcell.RuneVLine).
			Draw()
	}

	// Line intersection
	canvas.NewPoint(xVLine, yHLine).
		Foreground(tcell.ColorWhite).
		Char(tcell.RuneULCorner).
		Draw()

	// Cells
	for row := 0; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			bg, fg := tcell.ColorBlack, tcell.ColorSilver
			if row == grid.cursor.pos.row && col == grid.cursor.pos.col {
				bg, fg = fg, bg
			}
			cell := grid.GetCell(&Pos{row: row, col: col})
			x := xStart + lineWidth + col*cellWidth + rowIndexWidth
			y := yStart + lineHeight + row*cellHeight + cellHeight
			canvas.NewBox(x, y, cellWidth, cellHeight).
				Background(bg).
				Foreground(fg).
				Text(cell.value).
				PadRight(1).
				Draw()
		}
	}
}

func drawFooter() {
	canvas.NewBox(0, term.h-footerHeight, term.w, footerHeight).
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorSilver).
		Text("[No Name]").
		Draw()
}
