package main

import (
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

//------------------------------------------------------------------------------

type Swapper interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(r *rand.Rand, sw Swapper) {
	for n := sw.Len(); n > 0; n-- {
		sw.Swap(r.Intn(n), n-1)
	}
}

type Point2fSlice []Point2f

var _ Swapper = Point2fSlice(nil)

func (p Point2fSlice) Len() int      { return len(p) }
func (p Point2fSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p Point2fSlice) Shuffle(r *rand.Rand) {
	Shuffle(r, p)
}

//------------------------------------------------------------------------------

func randPoint2f(r *rand.Rand) Point2f {
	return randPoint2fMinMax(r, 0.05, 0.95)
}

func randPoint2fMinMax(r *rand.Rand, min, max float64) Point2f {
	return Point2f{
		X: lerp(min, max, r.Float64()),
		Y: lerp(min, max, r.Float64()),
	}
}
