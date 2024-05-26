package palgen

import (
	"image/color"
	"math"
)

type RGB struct {
	R, G, B float64
}

func (x RGB) Color() color.Color {
	return color.NRGBA64{
		R: colorComponentUint16(x.R),
		G: colorComponentUint16(x.G),
		B: colorComponentUint16(x.B),
		A: colorComponentUint16(1),
	}
}

func colorComponentUint16(c float64) uint16 {
	return uint16(math.Round(clamp01(c) * 0xffff))
}

func colorComponentUint32(c float64) uint32 {
	return uint32(colorComponentUint16(c))
}
