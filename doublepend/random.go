package main

import (
	"math"
	"math/rand"
	"time"
)

func newRandNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// [min..max)
func randIntInterval(r *rand.Rand, min, max int) int {
	if min >= max { // It is an empty interval.
		return 0
	}
	return min + r.Intn(max-min)
}

func randFloatInterval(r *rand.Rand, min, max float64) float64 {
	return lerp(min, max, r.Float64())
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

func randSamplesV1(r *rand.Rand, n int) []*Sample {
	dps := make([]*DoublePendulum, n)
	dp0 := randDoublePendulum(r)
	for i := 0; i < n; i++ {
		if i == 0 {
			dps[i] = dp0
		} else {
			clone := dp0.Clone()
			randChangeDoublePendulum(r, clone)
			dps[i] = clone
		}
	}
	samples := make([]*Sample, n)
	for i, dp := range dps {
		samples[i] = newSample(dp, GetPalette(i))
	}
	return samples
}

func randSamplesV2(r *rand.Rand, n int) []*Sample {

	k := 1.0
	dk := 0.5

	dMass := (massMax - massMin) / 100 // 1%

	dps := make([]*DoublePendulum, n)
	dp := randDoublePendulum(r)
	mass := dp[1].Mass
	for i := 0; i < n; i++ {
		clone := dp.Clone()
		p := &(clone[1])

		p.Mass = mass + dMass*k
		k *= dk

		dps[i] = clone
	}
	samples := make([]*Sample, n)
	for i, dp := range dps {
		samples[i] = newSample(dp, GetPalette(i))
	}
	return samples
}

func randSamplesV3(r *rand.Rand, n int) []*Sample {

	dMass := (massMax - massMin) / 100 // 1%
	dMass *= 0.001

	dps := make([]*DoublePendulum, n)
	dp := randDoublePendulum(r)
	mass := dp[1].Mass
	for i := 0; i < n; i++ {
		clone := dp.Clone()
		p := &(clone[1])

		p.Mass = mass
		mass += dMass

		dps[i] = clone
	}
	samples := make([]*Sample, n)
	for i, dp := range dps {
		samples[i] = newSample(dp, GetPalette(i))
	}
	return samples
}

func randSamples(r *rand.Rand, n int) []*Sample {
	//return randSamplesV1(r, n)
	//return randSamplesV2(r, n)
	return randSamplesV3(r, n)
}
