package main

import (
	"strconv"
)

// RowLabel returns the label for given row index
func RowLabel(row int) string {
	return strconv.Itoa(row + 1)
}

// ColLabel returns the label for given col index
// https://stackoverflow.com/a/182924
func ColLabel(col int) string {
	runes := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	runeLen := len(runes)
	label := ""
	n := col + 1
	var mod int
	for {
		if n <= 0 {
			break
		}
		mod = (n - 1) % runeLen
		label = string(runes[mod]) + label
		n = (n - mod) / runeLen
	}
	return label
}

// Cell represents a single cell in the grid
type Cell struct {
	row   int
	col   int
	value string
}

// Grid holds the state of the cell grid
type Grid struct {
	cells  map[int]map[int]*Cell
	cursor *Cursor
	// Total number of rows and columns
	rowTotal int
	colTotal int
	// Offset for pagination
	rowOff int
	colOff int
	// Limit for pagination
	rowLim int
	colLim int
}

const (
	defaultRowTotal = 1000
	defaultColTotal = 1000
)

var grid = &Grid{
	cells:    map[int]map[int]*Cell{},
	rowTotal: defaultRowTotal,
	colTotal: defaultColTotal,
}

func init() {
	// Set first cell as cursor
	cell := grid.GetCell(0, 0)
	grid.cursor = &Cursor{cell}
}

// GetCell retuns the cell at the position provided
// The cell will be created if it is not present
func (g *Grid) GetCell(row, col int) *Cell {
	var cellRow map[int]*Cell
	cellRow, ok := g.cells[row]
	if !ok {
		cellRow = map[int]*Cell{}
		g.cells[row] = cellRow
	}
	var cell *Cell
	cell, ok = cellRow[col]
	if !ok {
		cell = &Cell{row: row, col: col}
		cellRow[col] = cell
	}
	return cell
}

// SetCell sets the cell value for the position provided
// The cell will be created if it is not present
func (g *Grid) SetCell(row, col int, value string) {
	cell := g.GetCell(row, col)
	cell.value = value
}
