package main

// Cursor is the cell the cursor is on
type Cursor struct {
	*Cell
}

// MoveDown moves the grid cursor down
func (c *Cursor) MoveDown() {
	newPos := &Pos{row: c.pos.row + 1, col: c.pos.col}
	grid.cursor = &Cursor{grid.GetCell(newPos)}
}

// MoveUp moves the grid cursor up
func (c *Cursor) MoveUp() {
	if c.pos.row == 0 {
		return
	}
	newPos := &Pos{row: c.pos.row - 1, col: c.pos.col}
	grid.cursor = &Cursor{grid.GetCell(newPos)}
}

// MoveLeft moves the grid cursor left
func (c *Cursor) MoveLeft() {
	if c.pos.col == 0 {
		return
	}
	newPos := &Pos{row: c.pos.row, col: c.pos.col - 1}
	grid.cursor = &Cursor{grid.GetCell(newPos)}
}

// MoveRight moves the grid cursor right
func (c *Cursor) MoveRight() {
	newPos := &Pos{row: c.pos.row, col: c.pos.col + 1}
	grid.cursor = &Cursor{grid.GetCell(newPos)}
}
