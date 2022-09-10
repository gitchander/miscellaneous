package main

import (
	"fmt"
	"math"
	"math/rand"

	opt "github.com/gitchander/miscellaneous/attractor/utils/optional"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
)

func randTrig(r *rand.Rand) Trig {

	const (
		d = 1.0

		min = -math.Pi * d
		max = +math.Pi * d
	)

	return Trig{
		A: random.RandInterval(r, min, max),
		B: random.RandInterval(r, min, max),
		C: random.RandInterval(r, min, max),
		D: random.RandInterval(r, min, max),
	}
}

func randTrigSeed(seed int64) Trig {
	r := random.NewRandSeed(seed)
	return randTrig(r)
}

func randTrigOptSeed(optSeed opt.OptInt64) Trig {
	var seed int64
	if optSeed.Present {
		seed = optSeed.Value
	} else {
		r := random.NewRandNow()
		seed = r.Int63()
		fmt.Println("random seed:", seed)
	}
	return randTrigSeed(seed)
}
