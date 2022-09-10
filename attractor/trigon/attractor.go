package main

import (
	"math"

	"github.com/gitchander/miscellaneous/attractor"
	. "github.com/gitchander/miscellaneous/attractor/utils/point2f"
)

//------------------------------------------------------------------------------
// Trigonometric attractor
// x[n] = sin(a * y[n-1]) + cos(b * x[n-1])
// y[n] = sin(c * x[n-1]) + cos(d * y[n-1])

type Trig struct {
	A, B, C, D float64
}

var _ attractor.Feeder = Trig{}

func (t Trig) Feed(p Point2f) Point2f {
	return Point2f{
		X: math.Sin(t.A*p.Y) + math.Cos(t.B*p.X),
		Y: math.Sin(t.C*p.X) + math.Cos(t.D*p.Y),
	}
}

//------------------------------------------------------------------------------
