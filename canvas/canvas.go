package canvas

import (
	"github.com/gdamore/tcell/v2"
)

// Canvas is a wrapper for drawing on tcell.Screen
type Canvas struct {
	Screen tcell.Screen
}

// NewPoint creates a point that can be drawn
func (c *Canvas) NewPoint(x, y int) *Point {
	return &Point{
		canvas: c,
		X:      x,
		Y:      y,
	}
}

// NewBox creates a box that can be drawn
func (c *Canvas) NewBox(x, y, w, h int) *Box {
	return &Box{
		canvas: c,
		X:      x,
		Y:      y,
		W:      w,
		H:      h,
		Content: &Content{
			Pad: &Pad{},
		},
	}
}
