package drawer

// TextAlign represents text alignment options.
type TextAlign int

const (
	// AlignLeft aligns text to the left.
	AlignLeft TextAlign = iota

	// AlignCenter centers text.
	AlignCenter

	// AlignRight aligns text to the right.
	AlignRight
)

const (
	// FlexStart represents the start position for flex alignment (0.0).
	FlexStart = 0.0

	// FlexCenter represents the center position for flex alignment (0.5).
	FlexCenter = 0.5

	// FlexEnd represents the end position for flex alignment (1.0).
	FlexEnd = 1.0
)
