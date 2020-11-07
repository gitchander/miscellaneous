package main

import (
	"math"
)

const Tau = 2 * math.Pi

func DegToRad(deg float64) float64 {
	return deg * Tau / 360
}

func RadToDeg(rad float64) float64 {
	return rad * 360 / Tau
}

// func DegToRad(deg float64) float64 {
// 	return deg * math.Pi / 180
// }

// func RadToDeg(rad float64) float64 {
// 	return rad * 180 / math.Pi
// }

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

func angleNorm(angle float64) float64 {
	for angle < -math.Pi {
		angle += Tau
	}
	for angle > math.Pi {
		angle -= Tau
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

func mod(x, y int) int {
	m := x % y
	if m < 0 {
		m += y
	}
	return m
}
