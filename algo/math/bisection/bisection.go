package bisection

import (
	"errors"
	"math"
)

// Bisection method
// https://en.wikipedia.org/wiki/Bisection_method

func Bisection(f func(float64) float64, a, b float64,
	tolerance float64, nmax int) (float64, error) {

	if a >= b {
		return 0, errors.New("bisection: a >= b")
	}

	var (
		fa = f(a)
		fb = f(b)
	)

	if !differentSigns(fa, fb) {
		return 0, errors.New("bisection: sign(f(a)) = sign(f(b))")
	}

	for i := 0; i < nmax; i++ {

		c := (a + b) / 2 // midpoint of the interval [a,b]
		fc := f(c)

		if ((b - a) / 2) < tolerance {
			return c, nil
		}

		if differentSigns(fa, fc) {
			b = c
			fb = fc
		} else {
			a = c
			fa = fc
		}
	}

	return 0, errors.New("bisection: method failed")
}

var (
	// differentSigns = differentSignsV1
	// differentSigns = differentSignsV2
	differentSigns = differentSignsV3
)

func differentSignsV1(a, b float64) bool {
	return math.Signbit(a) != math.Signbit(b)
}

func differentSignsV2(a, b float64) bool {
	return (a * b) < 0
}

func differentSignsV3(a, b float64) bool {
	if (a < 0) && (0 < b) {
		return true
	}
	if (b < 0) && (0 < a) {
		return true
	}
	return false
}
