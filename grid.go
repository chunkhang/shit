package main

// Pos is the coordinate for a cell
type Pos struct {
	row int
	col int
}

// Grid holds the state of the cell grid
type Grid struct {
	row   int // Current row
	col   int // Current column
	cells map[int]map[int]string
}

var grid = &Grid{
	cells: map[int]map[int]string{},
}

func (g *Grid) hasValue(pos *Pos) bool {
	row, ok := g.cells[pos.row]
	if !ok {
		return false
	}
	value, ok := row[pos.col]
	return ok && value != ""
}

func (g *Grid) getValue(pos *Pos) string {
	if !g.hasValue(pos) {
		return ""
	}
	value, _ := g.cells[pos.row][pos.col]
	return value
}

func (g *Grid) setValue(pos *Pos, value string) {
	_, ok := g.cells[pos.row]
	if !ok {
		g.cells[pos.row] = map[int]string{}
	}
	g.cells[pos.row][pos.col] = value
}
