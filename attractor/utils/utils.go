package utils

import (
	"image/color"
	"math"
)

// Lerp - Linear interpolation
// t= [0, 1]
// (t == 0) => v0
// (t == 1) => v1
func Lerp(v0, v1 float64, t float64) float64 {
	return (1.0-t)*v0 + t*v1
}

func Round(a float64) int {
	//return int(math.Floor(a + 0.5))
	return int(math.Round(a))
}

func LerpColor(c0, c1 color.Color, t float64) color.Color {

	var (
		v0 = color.RGBAModel.Convert(c0).(color.RGBA)
		v1 = color.RGBAModel.Convert(c1).(color.RGBA)
	)

	var (
		r = uint8(Round(Lerp(float64(v0.R), float64(v1.R), t)))
		g = uint8(Round(Lerp(float64(v0.G), float64(v1.G), t)))
		b = uint8(Round(Lerp(float64(v0.B), float64(v1.B), t)))
		a = uint8(Round(Lerp(float64(v0.A), float64(v1.A), t)))
	)

	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
