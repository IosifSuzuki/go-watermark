package watermark

type Margins struct {
	Top    int
	Right  int
	Bottom int
	Left   int
}

func newMargins(horizontal, vertical int) Margins {
	return Margins{
		Top:    vertical,
		Right:  horizontal,
		Bottom: vertical,
		Left:   horizontal,
	}
}
