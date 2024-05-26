package palgen

import (
	"image"
)

func DrawPalette(m *image.RGBA, p Params) {

	var (
		bounds = m.Bounds()

		x0 = bounds.Min.X
		x1 = bounds.Max.X

		y0 = bounds.Min.Y
		y1 = bounds.Max.Y

		dx = bounds.Dx()
	)

	for x := x0; x < x1; x++ {
		t := float64(x-x0) / float64(dx) // [0..1)
		c := ColorByParams(p, t)
		for y := y0; y < y1; y++ {
			m.Set(x, y, c)
		}
	}
}
