package utils

import (
	"math"
)

// Lerp - Linear interpolation
// t= [0, 1]
// (t == 0) => v0
// (t == 1) => v1
func Lerp(v0, v1 float64, t float64) float64 {
	return (1.0-t)*v0 + t*v1
}

func Round(a float64) int {
	//return int(math.Floor(a + 0.5))
	return int(math.Round(a))
}
