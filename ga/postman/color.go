package main

import (
	"github.com/fogleman/gg"
)

type ColorRGB struct {
	R, G, B float64
}

func setColorRGB(dc *gg.Context, c ColorRGB) {
	dc.SetRGB(c.R, c.G, c.B)
}

var (
	clBlack = ColorRGB{R: 0, G: 0, B: 0}
	clWhite = ColorRGB{R: 1, G: 1, B: 1}

	clRed   = ColorRGB{R: 1, G: 0, B: 0}
	clGreen = ColorRGB{R: 0, G: 0.5, B: 0}
	clBlue  = ColorRGB{R: 0, G: 0, B: 1}
)
