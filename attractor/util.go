package attractor

import (
	"math"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func applyFactor(x, factor float64) float64 {
	if x < 0 {
		return 0
	}
	if x < 1 {
		return 1 - math.Exp(factor*math.Log(1-x))
	}
	return 1
}

func clamp(a float64, min, max float64) float64 {
	if a < min {
		a = min
	}
	if a > max {
		a = max
	}
	return a
}
