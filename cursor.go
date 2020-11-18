package main

// Cursor is the cell the cursor is on
type Cursor struct {
	*Cell
}

// MoveDown moves the grid cursor down
func (c *Cursor) MoveDown() {
	c.MoveTo(c.row+1, c.col)
}

// MoveUp moves the grid cursor up
func (c *Cursor) MoveUp() {
	c.MoveTo(c.row-1, c.col)
}

// MoveLeft moves the grid cursor left
func (c *Cursor) MoveLeft() {
	c.MoveTo(c.row, c.col-1)
}

// MoveRight moves the grid cursor right
func (c *Cursor) MoveRight() {
	c.MoveTo(c.row, c.col+1)
}

// MoveTo moves the grid cursor to the given position
func (c *Cursor) MoveTo(row, col int) {
	if row < 0 || col < 0 {
		return
	}
	if row >= grid.rowTotal || col >= grid.colTotal {
		return
	}

	// Update cursor
	grid.cursor = &Cursor{grid.GetCell(row, col)}

	// Update pagination
	if row <= grid.rowOff {
		grid.rowOff = Max(grid.rowOff-grid.rowLim, 0)
	}
	if row >= grid.rowOff+grid.rowLim {
		grid.rowOff += grid.rowLim
	}
	if col <= grid.colOff {
		grid.colOff = Max(grid.colOff-grid.colLim, 0)
	}
	if col >= grid.colOff+grid.colLim {
		grid.colOff += grid.colLim
	}
}
