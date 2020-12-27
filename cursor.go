package main

import (
	f "github.com/chunkhang/shit/file"
)

// Cursor is the cell the cursor is on
type Cursor struct {
	*f.Cell
}

var cursor *Cursor

// IsVisible checks whether the cursor is visible on screen
// After resizing the terminal, the cursor may not be visible
func (c *Cursor) IsVisible() bool {
	rowOK := c.Row >= grid.RowOff && c.Row < grid.RowOff+grid.RowLim
	colOK := c.Col >= grid.ColOff && c.Col < grid.ColOff+grid.ColLim
	return rowOK && colOK
}

// MoveDown moves the grid cursor down
func (c *Cursor) MoveDown() {
	c.MoveTo(c.Row+1, c.Col)
}

// MoveUp moves the grid cursor up
func (c *Cursor) MoveUp() {
	c.MoveTo(c.Row-1, c.Col)
}

// MoveLeft moves the grid cursor left
func (c *Cursor) MoveLeft() {
	c.MoveTo(c.Row, c.Col-1)
}

// MoveRight moves the grid cursor right
func (c *Cursor) MoveRight() {
	c.MoveTo(c.Row, c.Col+1)
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
	if row >= grid.RowTotal || col >= grid.ColTotal {
		return
	}

	// Update cursor
	cursor = &Cursor{grid.GetCell(row, col)}

	// Update pagination
	if row < grid.RowOff+scrollOff {
		grid.RowOff = Max(grid.RowOff-1, 0)
	}
	if row >= grid.RowOff+grid.RowLim-scrollOff {
		grid.RowOff++
	}
	if col < grid.ColOff+scrollOff {
		grid.ColOff = Max(grid.ColOff-1, 0)
	}
	if col >= grid.ColOff+grid.ColLim-scrollOff {
		grid.ColOff++
	}
}
