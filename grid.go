package main

import (
	"fmt"
)

// Pos is the coordinate for a cell
type Pos struct {
	row int
	col int
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
	cursor *Cursor
	cells  map[int]map[int]*Cell
}

var grid = &Grid{}

func init() {
	grid.cells = map[int]map[int]*Cell{}
	grid.cursor = &Cursor{grid.GetCell(&Pos{row: 0, col: 0})}
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
