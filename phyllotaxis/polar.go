package main

import "math"

// https://en.wikipedia.org/wiki/Polar_coordinate_system

// Rho - r, ρ
// Phi - φ, θ

type Polar struct {
	Rho float64
	Phi float64
}

func MakePolar(rho, phi float64) Polar {
	return Polar{
		Rho: rho,
		Phi: phi,
	}
}

func PolarToCartesian(p Polar) Point2f {
	sin, cos := math.Sincos(p.Phi)
	return MakePoint2f(cos, sin).MulScalar(p.Rho)
}

func CartesianToPolar(c Point2f) Polar {
	return Polar{
		Rho: math.Hypot(c.Y, c.X),
		Phi: math.Atan2(c.Y, c.X),
	}
}
