package canvas

import (
	"github.com/gdamore/tcell/v2"
)

// Point is the smallest drawable point
type Point struct {
	canvas *Canvas
	X      int
	Y      int
	Bg     tcell.Color
	Fg     tcell.Color
	Rev    bool
	Rune   rune
}

// Background sets the background for point
func (p *Point) Background(bg tcell.Color) *Point {
	p.Bg = bg
	return p
}

// Foreground sets the foreground for point
func (p *Point) Foreground(fg tcell.Color) *Point {
	p.Fg = fg
	return p
}

// Reverse reverses the background and foreground for point
func (p *Point) Reverse(on bool) *Point {
	p.Rev = on
	return p
}

// Char sets the rune for point
func (p *Point) Char(char rune) *Point {
	p.Rune = char
	return p
}

// Draw draws the point on canvas screen
func (p *Point) Draw() {
	style := tcell.StyleDefault.
		Background(p.Bg).
		Foreground(p.Fg).
		Reverse(p.Rev)
	p.canvas.Screen.SetContent(p.X, p.Y, p.Rune, nil, style)
}
