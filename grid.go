package main

import (
	"errors"
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

// Cursor is the cell the cursor is on
type Cursor struct {
	*Cell
}

func (c *Cursor) moveDown() {
	c.pos.row++
}

func (c *Cursor) moveUp() {
	if c.pos.row == 0 {
		return
	}
	c.pos.row--
}

func (c *Cursor) moveLeft() {
	if c.pos.col == 0 {
		return
	}
	c.pos.col--
}

func (c *Cursor) moveRight() {
	c.pos.col++
}

// Grid holds the state of the cell grid
type Grid struct {
	cursor *Cursor
	cells  map[int]map[int]*Cell
}

var (
	grid = &Grid{
		cursor: &Cursor{&Cell{pos: &Pos{row: 0, col: 0}}},
		cells:  map[int]map[int]*Cell{},
	}
	errNotFound = errors.New("not found")
)

func (g *Grid) getCell(pos *Pos) (cell *Cell, err error) {
	cellRow, ok := g.cells[pos.row]
	if !ok {
		return nil, errNotFound
	}
	cell, ok = cellRow[pos.col]
	if !ok {
		return nil, errNotFound
	}
	return
}

func (g *Grid) setCell(cell *Cell) {
	_, ok := g.cells[cell.pos.row]
	if !ok {
		g.cells[cell.pos.row] = map[int]*Cell{}
	}
	g.cells[cell.pos.row][cell.pos.col] = cell
}
