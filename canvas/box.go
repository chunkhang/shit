package canvas

import (
	"github.com/gdamore/tcell/v2"
)

// Box is a drawable rectangular area
type Box struct {
	canvas  *Canvas
	X       int
	Y       int
	W       int
	H       int
	Bg      tcell.Color
	Fg      tcell.Color
	Content *Content
}

// Background sets the background for box
func (b *Box) Background(bg tcell.Color) *Box {
	b.Bg = bg
	return b
}

// Foreground sets the text for box
func (b *Box) Foreground(fg tcell.Color) *Box {
	b.Fg = fg
	return b
}

// Text sets the text for box
func (b *Box) Text(text string) *Box {
	b.Content.Text = text
	return b
}

// AlignLeft sets the content alignment to left for box
func (b *Box) AlignLeft() *Box {
	b.Content.Align = AlignLeft
	return b
}

// AlignCenter sets the content alignment to center for box
func (b *Box) AlignCenter() *Box {
	b.Content.Align = AlignCenter
	return b
}

// AlignRight sets the content alignment to right for box
func (b *Box) AlignRight() *Box {
	b.Content.Align = AlignRight
	return b
}

// Pad sets the content left and right padding for box
func (b *Box) Pad(pad int) *Box {
	b.Content.Pad.Left = pad
	b.Content.Pad.Right = pad
	return b
}

// PadLeft sets the content left padding for box
func (b *Box) PadLeft(pad int) *Box {
	b.Content.Pad.Left = pad
	return b
}

// PadRight sets the content right padding for box
func (b *Box) PadRight(pad int) *Box {
	b.Content.Pad.Right = pad
	return b
}

// Draw draws the box on canvas screen
func (b *Box) Draw() {
	yStart := b.Y
	yEnd := b.Y + b.H

	xStart := b.X
	xEnd := b.X + b.W
	xLen := xEnd - xStart

	chars := []rune(b.Content.Text)
	charTotal := len(chars)
	i := 0

	for y := yStart; y < yEnd; y++ {
		// Calculate remaining characters for this row
		// This information is used to calculate the offset for alignment
		charOff := 0
		charLeft := charTotal - i
		if charLeft < xLen-b.Content.Pad.Left-b.Content.Pad.Right {
			if b.Content.Align == AlignRight {
				charOff = xLen - charLeft - b.Content.Pad.Right
			} else if b.Content.Align == AlignCenter {
				charOff = b.Content.Pad.Left + ((xLen - charLeft - b.Content.Pad.Right - b.Content.Pad.Left) / 2)
			}
		}
		for x := xStart; x < xEnd; x++ {
			point := b.canvas.NewPoint(x, y).Background(b.Bg).Foreground(b.Fg)
			hasChar := i < charTotal
			afterPadLeft := x-xStart >= b.Content.Pad.Left
			beforePadRight := xEnd-x > b.Content.Pad.Right
			afterOff := x-xStart >= charOff
			if hasChar && afterPadLeft && beforePadRight && afterOff {
				point.Char(chars[i])
				i++
			}
			point.Draw()
		}
	}
}
