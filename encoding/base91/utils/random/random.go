package random

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

func FillBytes(r *rand.Rand, bs []byte) {
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
