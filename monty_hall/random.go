package main

import (
	"math/rand"
	"time"
)

var getRand = func() func() *rand.Rand {
	rs := make(chan *rand.Rand)
	go func() {
		for {
			r := newRandNow()
			for i := 0; i < 1000; i++ {
				seed := r.Int63()
				rs <- newRandSeed(seed)
			}
		}
	}()
	return func() *rand.Rand {
		return <-rs
	}
}()

func newRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func newRandTime(t time.Time) *rand.Rand {
	return newRandSeed(t.UnixNano())
}

func newRandNow() *rand.Rand {
	return newRandTime(time.Now())
}
