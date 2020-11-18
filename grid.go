package main

import (
	"fmt"
	"strconv"
)

// Pos is the coordinate for a cell
type Pos struct {
	row int
	col int
}

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

// Label returns the location label for position
func (p *Pos) Label() string {
	return fmt.Sprintf("%s%s", ColLabel(p.col), RowLabel(p.row))
}

func (p *Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.row, p.col)
}

// Cell holds the state of a cell
type Cell struct {
	pos   *Pos
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

var grid = &Grid{
	cells:    map[int]map[int]*Cell{},
	rowTotal: 50,
	colTotal: 10,
}

func init() {
	// Set first cell as cursor
	cell := grid.GetCell(&Pos{row: 0, col: 0})
	grid.cursor = &Cursor{cell}
}

// GetCell retuns the cell at the position provided
// The cell will be created if it is not present
func (g *Grid) GetCell(pos *Pos) *Cell {
	var cellRow map[int]*Cell
	cellRow, ok := g.cells[pos.row]
	if !ok {
		cellRow = map[int]*Cell{}
		g.cells[pos.row] = cellRow
	}
	var cell *Cell
	cell, ok = cellRow[pos.col]
	if !ok {
		cell = &Cell{pos: pos}
		cellRow[pos.col] = cell
	}
	return cell
}
