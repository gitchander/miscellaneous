package core

import (
	"math/rand"
	"time"
)

func newRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func newRandTime(t time.Time) *rand.Rand {
	return newRandSeed(t.UnixNano())
}

func newRandNow() *rand.Rand {
	return newRandTime(time.Now())
}

func randFillBytes(r *rand.Rand, bs []byte) {
	const (
		bitsPerByte    = 8
		bytesPerUint64 = 8
	)
	var x uint64
	var n int // number of random bytes
	for i := range bs {
		if n == 0 {
			x = r.Uint64()
			n = bytesPerUint64
		}
		bs[i] = byte(x)
		x >>= bitsPerByte
		n--
	}
}

var nextRandom = func() func() *rand.Rand {
	c := make(chan *rand.Rand)
	go func() {
		r := newRandNow()
		for {
			c <- newRandSeed(r.Int63())
		}
	}()
	return func() *rand.Rand {
		return <-c
	}
}()
