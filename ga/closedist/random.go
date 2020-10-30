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

func randBool(r *rand.Rand) bool {
	return (r.Int() & 1) == 1
}

type Swapper interface {
	Len() int
	Swap(i, j int)
}

func shuffleElements(r *rand.Rand, sw Swapper) {
	n := sw.Len()
	for n > 0 {
		i := r.Intn(n)
		sw.Swap(i, n-1)
		n--
	}
}

type Point2fSlice []Point2f

var _ Swapper = Point2fSlice(nil)

func (p Point2fSlice) Len() int      { return len(p) }
func (p Point2fSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
