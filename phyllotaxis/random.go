package main

import (
	"image/color"
	"math/rand"
	"time"
)

func newRandNow() *rand.Rand {
	return newRandTime(time.Now())
}

func newRandTime(t time.Time) *rand.Rand {
	return newRandSeed(t.UnixNano())
}

func newRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

type Swapper interface {
	Len() int
	Swap(i, j int)
}

func shuffleElements(r *rand.Rand, sw Swapper) {
	for n := sw.Len(); n > 1; n-- {
		sw.Swap(r.Intn(n), n-1)
	}
}

type colorSlice []color.Color

var _ Swapper = colorSlice(nil)

func (p colorSlice) Len() int      { return len(p) }
func (p colorSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func shufflePalette(pal []color.Color) {
	r := newRandNow()
	shuffleElements(r, colorSlice(pal))
}
