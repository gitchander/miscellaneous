package main

import (
	"math"
)

const tau = 2 * math.Pi

const (
	goldAngleRad = tau / (math.Phi * math.Phi)
	goldAngleDeg = 360 / (math.Phi * math.Phi)
)

func DegToRad(deg float64) (rad float64) {
	return deg * (tau / 360)
}

func RadToDeg(rad float64) (deg float64) {
	return rad * (360 / tau)
}

// Vogel’s formula
// α - angle
func formulaV(n int, c, α float64) Polar {

	nf := float64(n)

	return Polar{
		Rho: c * math.Sqrt(nf),
		Phi: α * nf,
	}
}

func makePoints(n int) []Point2f {

	const angleRad = goldAngleRad

	center := MakePoint2f(1, 1).DivScalar(2)

	c := 0.05

	var ps []Point2f

	for i := 0; i < n; i++ {
		p := formulaV(i, c, angleRad)
		t := PolarToCartesian(p)
		t = t.Add(center)
		ps = append(ps, t)
	}

	return ps
}

func newIndividPhyllo(n int) *Individ {
	ps := makePoints(n)
	return NewIndivid(ps)
}
