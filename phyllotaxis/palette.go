package main

import (
	"image/color"
	"image/color/palette"

	"github.com/gitchander/miscellaneous/palgen"
	"github.com/gitchander/miscellaneous/utils/random"
)

func randPalette(n int, seed int64) []color.Color {
	//return randPalette1(n, seed)
	// return randPalette2(n, seed)
	return randPalette3(n, seed)
}

func randPalette1(n int, seed int64) []color.Color {
	r := random.NewRandSeed(seed)
	pal := clonePalette(palette.Plan9)
	shufflePalette(r, pal)
	return pal[:n]
}

func randPalette2(n int, seed int64) []color.Color {
	r := random.NewRandSeed(seed)
	pal := clonePalette(palette.WebSafe)
	shufflePalette(r, pal)
	return pal[:n]
}

func randPalette3(n int, seed int64) []color.Color {
	r := random.NewRandSeed(seed)
	var par palgen.Params
	palgen.RandParams(r, &par, n)
	pal := make([]color.Color, n)
	for i := range pal {
		pal[i] = palgen.ColorByParams(par, float64(i)/float64(n))
	}
	return pal
}

func clonePalette(a []color.Color) []color.Color {
	b := make([]color.Color, len(a))
	copy(b, a)
	return b
}
