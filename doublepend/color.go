package main

import (
	"github.com/gotk3/gotk3/cairo"
)

type ColorRGBf struct {
	R, G, B float64
}

func RGBf(r, g, b float64) ColorRGBf {
	return ColorRGBf{
		R: r,
		G: g,
		B: b,
	}.Normalize()
}

func (a ColorRGBf) Normalize() ColorRGBf {
	const (
		min = 0
		max = 1
	)
	return ColorRGBf{
		R: clampFloat64(a.R, min, max),
		G: clampFloat64(a.G, min, max),
		B: clampFloat64(a.B, min, max),
	}
}

func setColor(context *cairo.Context, a ColorRGBf) {
	context.SetSourceRGB(a.R, a.G, a.B)
	//context.SetSourceRGBA(a.R, a.G, a.B, 1)
}

func ColorLerp(c0, c1 ColorRGBf, t float64) ColorRGBf {
	return ColorRGBf{
		R: lerp(c0.R, c1.R, t),
		G: lerp(c0.G, c1.G, t),
		B: lerp(c0.B, c1.B, t),
	}.Normalize()
}

func Gray(x float64) ColorRGBf {
	return RGBf(x, x, x)
}

// Colors:
var (
	Black = RGBf(0, 0, 0)
	White = RGBf(1, 1, 1)

	Red  = RGBf(1, 0, 0)
	Lime = RGBf(0, 1, 0)
	Blue = RGBf(0, 0, 1)

	Green = RGBf(0, 0.5, 0) // hex: #008000, dec: (0,128,0)

	Yellow  = RGBf(1, 1, 0)
	Cyan    = RGBf(0, 1, 1)
	Magenta = RGBf(1, 0, 1)

	Gray25  = Gray(0.25)
	Gray50  = Gray(0.50)
	Gray75  = Gray(0.75)
	Gray100 = Gray(1.00)
)
