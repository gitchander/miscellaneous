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

func clerp(c0, c1 ColorRGBf, t float64) ColorRGBf {
	return ColorRGBf{
		R: lerp(c0.R, c1.R, t),
		G: lerp(c0.G, c1.G, t),
		B: lerp(c0.B, c1.B, t),
	}.Normalize()
}
