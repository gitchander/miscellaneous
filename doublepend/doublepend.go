package main

import (
	"math"
)

const gravity = 9.81

type DoublePendulum [2]Pendulum

func getDPCoords(dp DoublePendulum, scale float64) (x1, y1, x2, y2 float64) {

	var (
		p1 = &(dp[0])
		p2 = &(dp[1])
	)

	var (
		r1 = p1.Length * scale
		r2 = p2.Length * scale

		a1 = p1.Theta
		a2 = p2.Theta
	)

	sin1, cos1 := math.Sincos(a1)

	x1 = r1 * sin1
	y1 = r1 * cos1

	sin2, cos2 := math.Sincos(a2)

	x2 = x1 + r2*sin2
	y2 = y1 + r2*cos2

	return
}

func nextStep(dp *DoublePendulum, deltaTimeSec float64) {

	var (
		p1 = &(dp[0])
		p2 = &(dp[1])
	)

	p := &doublePendulum{
		m1: p1.Mass,
		m2: p2.Mass,

		r1: p1.Length,
		r2: p2.Length,

		a1: p1.Theta,
		a2: p2.Theta,

		v1: p1.Velocity,
		v2: p2.Velocity,
	}

	var dt = deltaTimeSec

	//---------------------------------------------------

	a1_a, a2_a := calculate(p)

	p.v1 += a1_a * dt
	p.v2 += a2_a * dt

	p.a1 += p.v1 * dt
	p.a2 += p.v2 * dt

	if false {
		const coef = 0.999
		p.v1 *= coef
		p.v2 *= coef
	}

	//---------------------------------------------------

	p1.Theta = p.a1
	p2.Theta = p.a2

	if true {
		p1.Theta = angleNormalize(p1.Theta)
		p2.Theta = angleNormalize(p2.Theta)
	}

	//---------------------------------------------------

	p1.Velocity = p.v1
	p2.Velocity = p.v2

	if true {
		p1.Velocity = clampVelocity(p1.Velocity)
		p2.Velocity = clampVelocity(p2.Velocity)
	}
}

type doublePendulum struct {
	m1, m2 float64 // Mass
	r1, r2 float64 // Length
	a1, a2 float64 // Theta
	v1, v2 float64 // Velocity
}

// type Thetas struct {
// 	theta    float64 // theta
// 	d_theta  float64 // theta'
// 	dd_theta float64 // theta"
// }

// accelerations: a1_a, a2_a
func calculate(p *doublePendulum) (a1_a, a2_a float64) {

	const g = gravity

	den := (2*p.m1 + p.m2 - (p.m2 * math.Cos(2*(p.a1-p.a2))))

	tmp1 := -g * ((2 * p.m1) + p.m2) * math.Sin(p.a1)
	tmp2 := -g * p.m2 * math.Sin(p.a1-2*p.a2)
	tmp3 := -2 * math.Sin(p.a1-p.a2) * p.m2
	tmp4 := p.v2*p.v2*p.r2 + (p.v1*p.v1)*(p.r1)*math.Cos(p.a1-p.a2)

	a1_a = (tmp1 + tmp2 + tmp3*tmp4) / (p.r1 * den)

	tmp1 = 2 * math.Sin(p.a1-p.a2)
	tmp2 = p.v1 * p.v1 * p.r1 * (p.m1 + p.m2)
	tmp3 = g * (p.m1 + p.m2) * math.Cos(p.a1)
	tmp4 = p.v2 * p.v2 * p.r2 * p.m2 * math.Cos(p.a1-p.a2)

	a2_a = (tmp1 * (tmp2 + tmp3 + tmp4)) / (p.r2 * den)

	return
}
