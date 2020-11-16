package canvas

// Align is the direction of text alignment
type Align int

const (
	// AlignLeft is left text alignment
	AlignLeft Align = iota
	// AlignCenter is center text alignment
	AlignCenter
	// AlignRight is right text alignment
	AlignRight
)

// Pad is the padding for text in box
type Pad struct {
	Left  int
	Right int
}

// Content is the content of a box
type Content struct {
	Text  string
	Align Align
	Pad   *Pad
}
