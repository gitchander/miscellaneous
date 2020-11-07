package main

import (
	"math/rand"
	"time"
)

func newRandNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randIntInterval(r *rand.Rand, min, max int) int {
	if min >= max { // It is an empty interval.
		return 0
	}
	return min + r.Intn(max-min)
}

func randPendulum(r *rand.Rand) Pendulum {

	angle := DegToRad(float64(randIntInterval(r, -180, 180)))

	return Pendulum{
		Mass:   float64(randIntInterval(r, 10, 50)),
		Length: float64(10 * randIntInterval(r, 8, 16)),
		Theta:  angle,
		//Velocity float64
	}
}

func randDoublePendulum(r *rand.Rand) *DoublePendulum {
	return &DoublePendulum{
		randPendulum(r),
		randPendulum(r),
	}
}

func randChangeDoublePendulum(r *rand.Rand, dp *DoublePendulum) {
	i := r.Intn(2)
	p := &(dp[i])

	n := r.Intn(3)
	switch n {
	case 0:
		p.Mass += 0.1 * r.Float64()
	case 1:
		p.Length += 0.1 * r.Float64()
	case 2:
		p.Theta += 0.1 * r.Float64()
	}
}
