package utils

import (
	"math/rand"
	"time"
)

func NewRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func NewRandTime(t time.Time) *rand.Rand {
	return NewRandSeed(t.UnixNano())
}

func NewRandNow() *rand.Rand {
	return NewRandTime(time.Now())
}

func RandIntMinMax(r *rand.Rand, min, max int) int {
	if min >= max {
		panic("invalid interval")
	}
	return min + r.Intn(max-min)
}
