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
	cells: map[int]map[int]string{
		0: map[int]string{
			0: "hello",
			2: "world",
		},
		4: map[int]string{
			3: "here",
			4: "!",
		},
	},
}

func (g *Grid) hasValue(pos *Pos) bool {
	row, ok := g.cells[pos.row]
	if !ok {
		return false
	}
	_, ok = row[pos.col]
	return ok
}

func (g *Grid) getValue(pos *Pos) string {
	if !g.hasValue(pos) {
		return ""
	}
	value, _ := g.cells[pos.row][pos.col]
	return value
}
