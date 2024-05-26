package random

import (
	"math/rand"
	"time"
)

func NewRandSeed(seed int64) *Rand {
	return rand.New(rand.NewSource(seed))
}

func NewRandTime(t time.Time) *Rand {
	return NewRandSeed(t.UnixNano())
}

func NewRandNow() *Rand {
	return NewRandTime(time.Now())
}
