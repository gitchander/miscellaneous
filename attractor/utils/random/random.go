package random

import (
	"math/rand"

	"github.com/gitchander/miscellaneous/attractor/utils"
)

func RandInterval(r *rand.Rand, min, max float64) float64 {
	t := r.Float64()
	return utils.Lerp(min, max, t)
}
