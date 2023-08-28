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
	var (
		x = p.Rho * cos
		y = p.Rho * sin
	)
	return MakePoint2f(x, y)
}

func CartesianToPolar(c Point2f) Polar {
	var (
		rho = math.Hypot(c.Y, c.X)
		phi = math.Atan2(c.Y, c.X)
	)
	return MakePolar(rho, phi)
}
