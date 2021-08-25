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

func randLengths(r *rand.Rand) (l1, l2 float64) {
	const (
		// totalLength = 1.0
		// minLength   = 0.25

		totalLength = 3.5
		minLength   = 1.0
	)
	for {
		l1 = r.Float64() * totalLength
		l2 = totalLength - l1
		if (l1 > minLength) && (l2 > minLength) {
			return
		}
	}
}

func randDoublePendulum(r *rand.Rand) DoublePendulum {

	l1, l2 := randLengths(r)

	var (
		p1 = Pendulum{
			Mass:   randMass(r),
			Length: l1,
			Theta:  randAngle(r),
		}
		p2 = Pendulum{
			Mass:   randMass(r),
			Length: l2,
			Theta:  randAngle(r),
		}
	)

	return DoublePendulum{p1, p2}
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
	massMin = 0.1
	massMax = 20.0
)

func randMass(r *rand.Rand) float64 {
	return randFloatInterval(r, massMin, massMax)
}

// const (
// 	lengthMin = 1.0
// 	lengthMax = 2.0
// )

// func randLength(r *rand.Rand) float64 {
// 	return randFloatInterval(r, lengthMin, lengthMax)
// }

const (
	angleMin = -math.Pi
	angleMax = +math.Pi
)

func randAngle(r *rand.Rand) float64 {
	return randFloatInterval(r, angleMin, angleMax)
}

func randSamplesV1(r *rand.Rand, n int) []*Sample {
	dps := make([]DoublePendulum, n)
	dp0 := randDoublePendulum(r)
	for i := 0; i < n; i++ {
		if i == 0 {
			dps[i] = dp0
		} else {
			clone := dp0
			randChangeDoublePendulum(r, &clone)
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

	dps := make([]DoublePendulum, n)
	dp := randDoublePendulum(r)
	mass := dp[1].Mass
	for i := 0; i < n; i++ {
		clone := dp
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

	dps := make([]DoublePendulum, n)
	dp := randDoublePendulum(r)
	mass := dp[1].Mass
	for i := 0; i < n; i++ {
		clone := dp
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
