package main

import (
	"math"
	"math/rand"

	attr "github.com/gitchander/miscellaneous/attractor"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
	"github.com/gitchander/miscellaneous/attractor/utils/random"
)

//------------------------------------------------------------------------------
// Trigonometric attractor
// x[n] = sin(a * y[n-1]) + cos(b * x[n-1])
// y[n] = sin(c * x[n-1]) + cos(d * y[n-1])

type Trigo struct {
	A, B, C, D float64
}

var _ attr.Feeder = &Trigo{}

func (t *Trigo) Feed(p Point2f) Point2f {
	return Point2f{
		X: math.Sin(t.A*p.Y) + math.Cos(t.B*p.X),
		Y: math.Sin(t.C*p.X) + math.Cos(t.D*p.Y),
	}.DivScalar(2)
}

func randTrigo(r *rand.Rand) *Trigo {

	const (
		d = 2.0 * math.Pi

		min = -d
		max = +d
	)

	return &Trigo{
		A: random.RandInterval(r, min, max),
		B: random.RandInterval(r, min, max),
		C: random.RandInterval(r, min, max),
		D: random.RandInterval(r, min, max),
	}
}
