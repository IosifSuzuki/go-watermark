package watermark

import (
	"image"
	"image/color"
	"math"
)

func newAlphaMask(opacity float64) *image.Uniform {
	opacity = math.Max(0, math.Min(1, opacity))
	return image.NewUniform(color.Alpha{A: uint8(opacity * 255)})
}
