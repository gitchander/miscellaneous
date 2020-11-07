package main

import (
	"math"
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

func randFloatInterval(r *rand.Rand, min, max float64) float64 {
	if min > max { // It is an empty interval.
		return 0
	}
	t := r.Float64()
	return lerp(min, max, t)
}

func randPendulum(r *rand.Rand) Pendulum {
	return Pendulum{
		Mass:   randMass(r),
		Length: randLength(r),
		Theta:  randAngle(r),
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

	// n := r.Intn(3)
	// switch n {
	// case 0:
	// 	p.Mass += 0.1 * r.Float64()
	// case 1:
	// 	p.Length += 0.1 * r.Float64()
	// case 2:
	// 	p.Theta += 0.1 * r.Float64()
	// }

	p.Mass += 0.001 * r.Float64()
}

const (
	massMin = 0.2
	massMax = 20.0
)

func randMass(r *rand.Rand) float64 {
	return randFloatInterval(r, massMin, massMax)
}

const (
	lengthMin = 1.0
	lengthMax = 2.0
)

const lengthScale = 100.0

func randLength(r *rand.Rand) float64 {
	return randFloatInterval(r, lengthMin, lengthMax)
}

const (
	angleMin = -math.Pi
	angleMax = +math.Pi
)

func randAngle(r *rand.Rand) float64 {
	return randFloatInterval(r, angleMin, angleMax)
}
