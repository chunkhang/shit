package main

// Cursor is the cell the cursor is on
type Cursor struct {
	*Cell
}

// IsVisible checks whether the cursor is visible on screen
// After resizing the terminal, the cursor may not be visible
func (c *Cursor) IsVisible() bool {
	rowOK := c.row >= grid.rowOff && c.row < grid.rowOff+grid.rowLim
	colOK := c.col >= grid.colOff && c.col < grid.colOff+grid.colLim
	return rowOK && colOK
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

const (
	// Scroll offset of 0 provides the smoothest scrolling
	scrollOff = 0
)

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
	if row < grid.rowOff+scrollOff {
		grid.rowOff = Max(grid.rowOff-1, 0)
	}
	if row >= grid.rowOff+grid.rowLim-scrollOff {
		grid.rowOff++
	}
	if col < grid.colOff+scrollOff {
		grid.colOff = Max(grid.colOff-1, 0)
	}
	if col >= grid.colOff+grid.colLim-scrollOff {
		grid.colOff++
	}
}
