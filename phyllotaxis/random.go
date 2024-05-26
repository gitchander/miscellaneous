package main

import (
	"image/color"

	"github.com/gitchander/miscellaneous/utils/random"
)

type colorSlice []color.Color

var _ random.Swapper = colorSlice(nil)

func (p colorSlice) Len() int      { return len(p) }
func (p colorSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func shufflePalette(r *random.Rand, pal []color.Color) {
	random.Shuffle(r, colorSlice(pal))
}
