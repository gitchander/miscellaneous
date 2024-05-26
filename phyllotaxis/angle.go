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
