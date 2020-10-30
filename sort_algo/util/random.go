package util

import (
	"math/rand"
	"time"
)

type Randomer interface {
	Intn(n int) int // 0 <= x < n
}

func RandomerSeed(seed int64) Randomer {
	return rand.New(rand.NewSource(seed))
}

func RandomerTime(t time.Time) Randomer {
	return rand.New(rand.NewSource(t.UnixNano()))
}

func RandomerTimeNow() Randomer {
	return rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}
