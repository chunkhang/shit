package file

// Grid holds the state of the cell grid
type Grid struct {
	cells map[int]map[int]*Cell
	// Total number of rows and columns
	RowTotal int
	ColTotal int
	// Offset for pagination
	RowOff int
	ColOff int
	// Limit for pagination
	RowLim int
	ColLim int
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
		cell = &Cell{Row: row, Col: col}
		cellRow[col] = cell
	}
	return cell
}

// SetCell sets the cell value for the position provided
// The cell will be created if it is not present
func (g *Grid) SetCell(row, col int, value string) {
	cell := g.GetCell(row, col)
	cell.Value = value
}
