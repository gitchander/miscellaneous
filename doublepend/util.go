package main

import (
	"math"
)

func DegToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func clampFloat64(a float64, min, max float64) float64 {
	if max < min { // empty range
		return 0
	}
	if a < min {
		a = min
	}
	if a > max {
		a = max
	}
	return a
}

func ceilPowerOfTwo(x int) int {
	d := 1
	for d < x {
		d *= 2
	}
	return d
}

const twoPi = 2 * math.Pi

func angleNorm(angle float64) float64 {
	for angle < -math.Pi {
		angle += twoPi
	}
	for angle > math.Pi {
		angle -= twoPi
	}
	return angle
}

// Lerp - Linear interpolation
// t= [0..1]
// (t == 0) => v0
// (t == 1) => v1
func lerp(v0, v1 float64, t float64) float64 {
	return (1.0-t)*v0 + t*v1
}
