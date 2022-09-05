package main

import (
	"fmt"
	"math"
	"math/rand"

	opt "github.com/gitchander/miscellaneous/attractor/utils/optional"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
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

func randFirstPoint() Point2f {
	const d = 2
	var (
		min = Pt2f(-d, -d)
		max = Pt2f(d, d)
	)
	r := random.NewRandNow()
	return Point2f{
		X: random.RandInterval(r, min.X, max.X),
		Y: random.RandInterval(r, min.Y, max.Y),
	}
}
