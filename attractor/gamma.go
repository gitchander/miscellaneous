package attractor

import (
	"math"
)

// Gamma correction
// https://en.wikipedia.org/wiki/Gamma_correction

type CorrectFunc func(inputValue float64) (outputValue float64)

// gamma < 1 - encoding gamma
// gamma > 1 - decoding gamma

// type GammaCorrection struct {
// 	Gamma float64 `json:"gamma"`
// }

func GammaCorrectionFunc(gamma float64) CorrectFunc {
	if gamma == 1 {
		// linear
		return func(inputValue float64) (outputValue float64) {
			outputValue = inputValue
			return
		}
	}
	return func(inputValue float64) (outputValue float64) {
		outputValue = math.Pow(inputValue, gamma)
		return
	}
}

// shade
func ToneCorrectionFunc(toneFactor float64) CorrectFunc {
	factor := toneFactor
	return func(x float64) float64 {
		if x < 0 {
			return 0
		}
		if x < 1 {
			return 1 - math.Exp(factor*math.Log(1-x))
		}
		return 1
	}
}

// 1 - (1 - x) ^ factor
// 1 - exp(factor * log(1-x))
func applyFactor(x, factor float64) float64 {
	if x < 0 {
		return 0
	}
	if x < 1 {
		return 1 - math.Exp(factor*math.Log(1-x))
	}
	return 1
}
